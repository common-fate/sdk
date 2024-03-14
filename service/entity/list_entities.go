package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/patrickmn/go-cache"
)

type ListInput struct {
	Type string
	ListOptions
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

type ListOptions struct {
	PageToken       string
	IncludeArchived bool
	OrderDescending bool
}

// All retrieves all entities of a specified type from the data source.
//
// Parameters:
//
//	ctx: The context.Context for the request.
//	c: The *Client object used to make the request.
//	input: A struct containing options for listing entities, such as pagination, inclusion of archived entities,
//	       and sorting order.
//
// Returns:
//
//	[]T: A slice of entities of type T retrieved from the data source.
//	error: An error, if any, encountered during the retrieval process.
//
// Description:
//
//	All is a generic function designed to retrieve all entities of a specified type from a data source using a
//	provided *Client object. The function iterates over pages of entities and unmarshals each entity into the
//	specified type T, accumulating them into a slice returned to the caller. It handles pagination, error handling,
//	and other aspects of the retrieval process internally.
//
// Example Usage:
//
//	entities, err := All[cf.User](ctx, client, ListOptions{
//	    PageToken:       "some_page_token",
//	    IncludeArchived: true,
//	    OrderDescending: false,
//	})
//	if err != nil {
//	    // Handle error
//	}
//	// entities will be a slice of type cf.User
func All[T Entity](ctx context.Context, c *Client, input ListOptions) ([]T, error) {
	e := *new(T)
	list := c.ListRequest(ListInput{
		Type:        e.EID().Type,
		ListOptions: input,
	})
	var out []T
	err := list.Pages(ctx, func(lo *ListOutput) error {
		for _, ent := range lo.Entities {
			var outType T
			err := Unmarshal(ent, outType)
			if err != nil {
				return err
			}
			out = append(out, outType)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}
