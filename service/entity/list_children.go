package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"go.opentelemetry.io/otel/attribute"
)

type ListChildrenInput struct {
	Parent    eid.EID
	PageToken string
}

type ListChildrenOutput struct {
	Children      []eid.EID
	NextPageToken string
}

func (c *Client) ListChildren(ctx context.Context, input ListChildrenInput) (ListChildrenOutput, error) {
	ctx, span := tracer.Start(ctx, "ListChildren")
	defer span.End()

	req := &entityv1alpha1.ListChildrenRequest{
		Universe:  "default",
		PageToken: input.PageToken,
		Parent:    input.Parent.ToAPI(),
	}

	res, err := c.raw.ListChildren(ctx, connect.NewRequest(req))
	if err != nil {
		return ListChildrenOutput{}, err
	}

	out := ListChildrenOutput{
		NextPageToken: res.Msg.NextPageToken,
	}

	for _, p := range res.Msg.Children {
		out.Children = append(out.Children, eid.FromAPI(p))
	}

	span.SetAttributes(attribute.Int("children_count", len(res.Msg.Children)))
	span.SetAttributes(attribute.String("parent", input.Parent.ID))

	return out, nil
}

type listChildrenRequestCall struct {
	input  ListChildrenInput
	client *Client
}

func (c *Client) ListChildrenRequest(input ListChildrenInput) *listChildrenRequestCall {
	return &listChildrenRequestCall{
		input:  input,
		client: c,
	}
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *listChildrenRequestCall) Pages(ctx context.Context, f func(ListChildrenOutput) error) error {
	// resets the input back to its original state
	originalPageToken := c.input.PageToken
	defer func() { c.input.PageToken = originalPageToken }()
	for {
		x, err := c.client.ListChildren(ctx, c.input)
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
