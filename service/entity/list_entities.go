package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/patrickmn/go-cache"
)

type ListInput struct {
	Type            string
	PageToken       string
	IncludeArchived bool
	OrderDescending bool
}

func (c *Client) List(ctx context.Context, input ListInput) (*ListOutput, error) {
	req := &entityv1alpha1.ListRequest{
		Universe:        "default",
		Type:            input.Type,
		PageToken:       input.PageToken,
		IncludeArchived: input.IncludeArchived,
	}
	if input.OrderDescending {
		req.Order = entityv1alpha1.Order_ORDER_DESCENDING.Enum().Enum()
	}

	res, err := c.raw.List(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	// update the attribute cache
	for _, e := range res.Msg.Entities {
		id := eid.EID{
			Type: e.Eid.Type,
			ID:   e.Eid.Id,
		}

		c.cache.Set(id.String(), e, cache.DefaultExpiration)
	}

	return res.Msg, nil
}

type listEntitiesRequestCall struct {
	input  ListInput
	client *Client
}

// ListRequest returns a request with listEntitiesRequestCall.Pages() that will pull all pages of results, invoking the callback for each page
func (c *Client) ListRequest(input ListInput) *listEntitiesRequestCall {
	return &listEntitiesRequestCall{
		input:  input,
		client: c,
	}
}

type ListOutput = entityv1alpha1.ListResponse

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *listEntitiesRequestCall) Pages(ctx context.Context, f func(*ListOutput) error) error {
	// resets the input back to its original state
	originalPageToken := c.input.PageToken
	defer func() { c.input.PageToken = originalPageToken }()
	for {
		x, err := c.client.List(ctx, c.input)
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.input.PageToken = x.NextPageToken
	}
}

// All returns all results from a paginated API
func (c *Client) All(ctx context.Context, input ListInput) (out []*entityv1alpha1.Entity, err error) {
	err = c.ListRequest(input).Pages(ctx, func(lo *ListOutput) error {
		out = append(out, lo.Entities...)
		return nil
	})
	return
}
