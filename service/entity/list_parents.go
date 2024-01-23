package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
)

type ListParentsInput struct {
	Child     eid.EID
	PageToken string
}

type ListParentsOutput struct {
	Parents       []eid.EID
	NextPageToken string
}

func (c *Client) ListParents(ctx context.Context, input ListParentsInput) (ListParentsOutput, error) {
	req := &entityv1alpha1.ListParentsRequest{
		Universe:  "default",
		PageToken: input.PageToken,
		Child:     input.Child.ToAPI(),
	}

	res, err := c.raw.ListParents(ctx, connect.NewRequest(req))
	if err != nil {
		return ListParentsOutput{}, err
	}

	out := ListParentsOutput{
		NextPageToken: res.Msg.NextPageToken,
	}

	for _, p := range res.Msg.Parents {
		out.Parents = append(out.Parents, eid.FromAPI(p))
	}

	return out, nil
}

type listParentsRequestCall struct {
	input  ListParentsInput
	client *Client
}

func (c *Client) ListParentsRequest(input ListParentsInput) *listParentsRequestCall {
	return &listParentsRequestCall{
		input:  input,
		client: c,
	}
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *listParentsRequestCall) Pages(ctx context.Context, f func(ListParentsOutput) error) error {
	// resets the input back to its original state
	originalPageToken := c.input.PageToken
	defer func() { c.input.PageToken = originalPageToken }()
	for {
		x, err := c.client.ListParents(ctx, c.input)
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
