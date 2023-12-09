package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/patrickmn/go-cache"
)

type ListInput struct {
	Type            string
	PageToken       string
	IncludeArchived bool
}

func (c *Client) List(ctx context.Context, input ListInput) (*ListOutput, error) {
	req := &entityv1alpha1.ListRequest{
		Universe:        "default",
		Type:            input.Type,
		PageToken:       input.PageToken,
		IncludeArchived: input.IncludeArchived,
	}

	res, err := c.raw.List(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	// update the attribute cache
	for _, e := range res.Msg.Entities {
		id := uid.UID{
			Type: e.Uid.Type,
			ID:   e.Uid.Id,
		}

		c.cache.Set(id.String(), e, cache.DefaultExpiration)
	}

	return res.Msg, nil
}

type listEntitiesRequestCall struct {
	input  ListInput
	client *Client
}

// FilterEntitiesRequest returns a request with filterEntitiesRequestCall.Pages() that will pull all pages of results, invoking the callback for each page
// I based this pattern off the google cloud SDK, I found it to be pretty neat, not set on the naming
// I think a good API here will have the option to do a single API call or a Pages call
// in the google API it would be filterEntitiesRequestCall.Do() to make a single request
// they also use a chained builder pattern
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
