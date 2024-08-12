package ingest

import "time"

type Event interface {
	EventType() string
}

type AWSAssumeRoleEvent struct {
	AccountID   string
	RoleName    string
	AccessKeyID string
	Timestamp   time.Time
}

func (AWSAssumeRoleEvent) EventType() string {
	return "AWSAssumeRoleEvent"
}

type AWSAPICallAttemptEvent struct {
	ClientID       string
	Service        string
	API            string
	AttemptLatency uint32
	FQDN           string
	UserAgent      string
	AccessKey      string
	Region         string
	HTTPStatusCode uint32
	XAmzRequestID  string
	XAmzID2        string
	Timestamp      time.Time
	Parameters     map[string][]string
	URIParameters  map[string][]string
}

func (AWSAPICallAttemptEvent) EventType() string {
	return "AWSAPICallAttemptEvent"
}

type AWSAPICallEvent struct {
	ClientID            string
	Service             string
	API                 string
	AttemptCount        uint32
	Region              string
	UserAgent           string
	FinalHTTPStatusCode uint32
	AttemptLatency      uint32
	MaxRetriesExceeded  uint32
	Timestamp           time.Time
	Parameters          map[string][]string
	URIParameters       map[string][]string
}

func (AWSAPICallEvent) EventType() string {
	return "AWSAPICallEvent"
}
