package policyset

import (
	"context"

	"connectrpc.com/connect"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type ListInput struct {
	PageToken string
}

type ListOutput struct {
	PolicySets    []PolicySet `json:"policySets"`
	NextPageToken string      `json:"nextPageToken,omitempty"`
}

func (c *Client) List(ctx context.Context, input ListInput) (ListOutput, error) {
	apires, err := c.raw.ListPolicySets(ctx, connect.NewRequest(&authzv1alpha1.ListPolicySetsRequest{
		PageToken: input.PageToken,
	}))
	if err != nil {
		return ListOutput{}, err
	}

	res := ListOutput{
		NextPageToken: apires.Msg.NextPageToken,
		PolicySets:    make([]PolicySet, len(apires.Msg.PolicySets)),
	}

	for i, p := range apires.Msg.PolicySets {
		res.PolicySets[i] = FromAPI(p)
	}

	return res, nil
}

type listPolicySetsRequestCall struct {
	input  ListInput
	client *Client
}

func (c *Client) ListPolicySetsRequest(input ListInput) *listPolicySetsRequestCall {
	return &listPolicySetsRequestCall{
		input:  input,
		client: c,
	}
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *listPolicySetsRequestCall) Pages(ctx context.Context, f func(ListOutput) error) error {
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
