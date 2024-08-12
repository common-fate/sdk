package ingest

import (
	"context"
	"errors"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/config"
	ingestv1alpha1 "github.com/common-fate/sdk/gen/granted/ingest/v1alpha1"
	"github.com/common-fate/sdk/gen/granted/ingest/v1alpha1/ingestv1alpha1connect"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrClosed          = errors.New("the client was already closed")
	ErrTooManyRequests = errors.New("too many requests are already in-flight")
)

type Client struct {
	apiClient ingestv1alpha1connect.IngestServiceClient

	msgs chan Event

	// These two channels are used to synchronize the client shutting down when
	// `Close` is called.
	// The first channel is closed to signal the backend goroutine that it has
	// to stop, then the second one is closed by the backend goroutine to signal
	// that it has finished flushing all queued messages.
	quit     chan struct{}
	shutdown chan struct{}

	Interval  time.Duration
	BatchSize int

	Logger *zap.SugaredLogger

	maxConcurrentRequests int

	// OnSendFailure is an optional but recommended callback function
	// which is called when there is a failure sending events.
	OnSendFailure func(event Event, err error)

	// OnSendSuccess is an optional callback function
	// which is called when events are successfully sent
	OnSendSuccess func(event Event)

	WriteContext context.Context
}

// NewFromConfig creates a new ingest client and starts a background goroutine
// to dispatch batch messages.
//
// Example usage:
//
//	client := ingest.NewFromConfig(cfg)
//	defer client.Close()
func NewFromConfig(cfg *config.Context) *Client {
	apiClient := ingestv1alpha1connect.NewIngestServiceClient(cfg.HTTPClient, cfg.APIURL)

	c := &Client{
		Logger:                zap.S(),
		msgs:                  make(chan Event, 100),
		quit:                  make(chan struct{}),
		shutdown:              make(chan struct{}),
		apiClient:             apiClient,
		Interval:              time.Second,
		maxConcurrentRequests: 100,
		BatchSize:             50,
		WriteContext:          context.Background(),
	}
	go c.loop()

	return c
}

func (c *Client) Track(event Event) {
	c.msgs <- event
}

// Batch loop.
func (c *Client) loop() {
	defer close(c.shutdown)

	wg := &sync.WaitGroup{}
	defer wg.Wait()

	tick := time.NewTicker(c.Interval)
	defer tick.Stop()

	ex := newExecutor(c.maxConcurrentRequests)
	defer ex.close()

	mq := messageQueue{
		maxBatchSize: c.BatchSize,
	}

	for {
		select {
		case msg := <-c.msgs:
			c.push(&mq, msg, wg, ex)

		case <-tick.C:
			c.flush(&mq, wg, ex)

		case <-c.quit:
			c.Logger.Debugf("exit requested – draining messages")

			// Drain the msg channel, we have to close it first so no more
			// messages can be pushed and otherwise the loop would never end.
			close(c.msgs)
			for msg := range c.msgs {
				c.push(&mq, msg, wg, ex)
			}

			c.flush(&mq, wg, ex)
			c.Logger.Debugf("exit")
			return
		}
	}
}

func (c *Client) push(q *messageQueue, m Event, wg *sync.WaitGroup, ex *executor) {
	c.Logger.Debugf("buffer (%d/%d) %v", len(q.pending), c.BatchSize, m)

	if msgs := q.push(m); msgs != nil {
		c.Logger.Debugf("exceeded messages batch limit with batch of %d messages – flushing", len(msgs))
		c.sendAsync(msgs, wg, ex)
	}
}

func (c *Client) flush(q *messageQueue, wg *sync.WaitGroup, ex *executor) {
	if msgs := q.flush(); msgs != nil {
		c.Logger.Debugf("flushing %d messages", len(msgs))
		c.sendAsync(msgs, wg, ex)
	}
}

// Close and flush events.
func (c *Client) Close() (err error) {
	defer func() {
		// Always recover, a panic could be raised if `c`.quit was closed which
		// means the method was called more than once.
		if recover() != nil {
			err = ErrClosed
		}
	}()
	close(c.quit)
	<-c.shutdown
	return
}

// Asynchronously send a batched requests.
func (c *Client) sendAsync(msgs []Event, wg *sync.WaitGroup, ex *executor) {
	wg.Add(1)

	if !ex.do(func() {
		defer wg.Done()
		defer func() {
			// In case a bug is introduced in the send function that triggers
			// a panic, we don't want this to ever crash the application so we
			// catch it here and log it instead.
			if err := recover(); err != nil {
				c.Logger.Errorf("panic - %s", err)
			}
		}()
		c.send(msgs)
	}) {
		wg.Done()
		c.Logger.Errorf("sending messages failed - %s", ErrTooManyRequests)
		c.notifyFailure(msgs, ErrTooManyRequests)
	}
}

// Send batch request.
func (c *Client) send(msgs []Event) {
	req := ingestv1alpha1.BatchWriteEventsRequest{
		Events: make([]*ingestv1alpha1.Event, len(msgs)),
	}

	for i, msg := range msgs {
		switch t := msg.(type) {
		case AWSAssumeRoleEvent:
			req.Events[i] = &ingestv1alpha1.Event{
				Timestamp: timestamppb.New(t.Timestamp),
				Details: &ingestv1alpha1.Event_AwsAssumeRole{
					AwsAssumeRole: &ingestv1alpha1.AWSAssumeRoleEvent{
						AccountId:   t.AccountID,
						RoleName:    t.RoleName,
						AccessKeyId: t.AccessKeyID,
					},
				},
			}
		case AWSAPICallAttemptEvent:
			req.Events[i] = &ingestv1alpha1.Event{
				Timestamp: timestamppb.New(t.Timestamp),
				Details: &ingestv1alpha1.Event_AwsApiCallAttempt{
					AwsApiCallAttempt: &ingestv1alpha1.AWSAPICallAttemptEvent{
						ClientId:       t.ClientID,
						AttemptLatency: t.AttemptLatency,
						Fqdn:           t.FQDN,
						UserAgent:      t.UserAgent,
						AccessKey:      t.AccessKey,
						Region:         t.Region,
						HttpStatusCode: t.HTTPStatusCode,
						XAmzRequestId:  t.XAmzRequestID,
						XAmzId_2:       t.XAmzID2,
						Service:        t.Service,
						Api:            t.API,
						Parameters:     convertParams(t.Parameters),
						UriParameters:  convertURIParams(t.URIParameters),
					},
				},
			}
		case AWSAPICallEvent:
			req.Events[i] = &ingestv1alpha1.Event{
				Timestamp: timestamppb.New(t.Timestamp),
				Details: &ingestv1alpha1.Event_AwsApiCall{
					AwsApiCall: &ingestv1alpha1.AWSAPICallEvent{
						ClientId:            t.ClientID,
						AttemptLatency:      t.AttemptLatency,
						UserAgent:           t.UserAgent,
						Region:              t.Region,
						Service:             t.Service,
						Api:                 t.API,
						AttemptCount:        t.AttemptCount,
						FinalHttpStatusCode: t.FinalHTTPStatusCode,
						MaxRetriesExceeded:  t.MaxRetriesExceeded,
						Parameters:          convertParams(t.Parameters),
						UriParameters:       convertURIParams(t.URIParameters),
					},
				},
			}
		}
	}

	_, err := c.apiClient.BatchWriteEvents(c.WriteContext, connect.NewRequest(&req))
	if err != nil {
		c.notifyFailure(msgs, err)
		return
	}

	c.notifySuccess(msgs)
}

func convertParams(input map[string][]string) []*ingestv1alpha1.Parameter {
	params := make([]*ingestv1alpha1.Parameter, 0, len(input))
	for k, v := range input {
		params = append(params, &ingestv1alpha1.Parameter{
			Key:   k,
			Value: v,
		})
	}

	return params
}

func convertURIParams(input map[string]string) []*ingestv1alpha1.Parameter {
	params := make([]*ingestv1alpha1.Parameter, 0, len(input))
	for k, v := range input {
		params = append(params, &ingestv1alpha1.Parameter{
			Key:   k,
			Value: []string{v},
		})
	}

	return params
}

func (c *Client) notifyFailure(msgs []Event, err error) {
	if c.OnSendFailure != nil {
		for _, m := range msgs {
			c.OnSendFailure(m, err)
		}
	}
}

func (c *Client) notifySuccess(msgs []Event) {
	if c.OnSendSuccess != nil {
		for _, m := range msgs {
			c.OnSendSuccess(m)
		}
	}
}
