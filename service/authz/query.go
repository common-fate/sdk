package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/patrickmn/go-cache"
)

type QueryInput struct {
	Type            string
	DirectParents   []uid.UID
	AttributeEquals []Attribute
	PageToken       string
}

type Attribute interface {
	ToAPI() *authzv1alpha1.Attribute
}

func (c *Client) Query(ctx context.Context, input QueryInput) (*authzv1alpha1.QueryResponse, error) {
	req := &authzv1alpha1.QueryRequest{
		Universe:        "default",
		Type:            input.Type,
		DirectParents:   make([]*authzv1alpha1.UID, len(input.DirectParents)),
		AttributeEquals: make([]*authzv1alpha1.Attribute, len(input.AttributeEquals)),
		PageToken:       input.PageToken,
	}

	for i, a := range input.AttributeEquals {
		req.AttributeEquals[i] = a.ToAPI()
	}

	for i, p := range input.DirectParents {
		req.DirectParents[i] = p.ToAPI()
	}

	res, err := c.raw.Query(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	// update the attribute cache
	for _, e := range res.Msg.Entities {
		id := uid.UID{
			Type: e.Entity.Uid.Type,
			ID:   e.Entity.Uid.Id,
		}

		c.cache.Set(id.String(), e, cache.DefaultExpiration)
	}

	return res.Msg, nil
}

type queryRequestCall struct {
	input  QueryInput
	client *Client
}

// FilterEntitiesRequest returns a request with filterEntitiesRequestCall.Pages() that will pull all pages of results, invoking the callback for each page
// I based this pattern off the google cloud SDK, I found it to be pretty neat, not set on the naming
// I think a good API here will have the option to do a single API call or a Pages call
// in the google API it would be filterEntitiesRequestCall.Do() to make a single request
// they also use a chained builder pattern
func (c *Client) QueryRequest(input QueryInput) *queryRequestCall {
	return &queryRequestCall{
		input:  input,
		client: c,
	}
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *queryRequestCall) Pages(ctx context.Context, f func(*authzv1alpha1.QueryResponse) error) error {
	// resets the input back to its original state
	originalPageToken := c.input.PageToken
	defer func() { c.input.PageToken = originalPageToken }()
	for {
		x, err := c.client.Query(ctx, c.input)
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
