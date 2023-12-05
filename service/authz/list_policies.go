package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type ListPolicySetsInput struct {
	PageToken string
}

type ListPolicySetsOutput struct {
	PolicySets    []PolicySet
	NextPageToken string
}

func (c *Client) ListPolicySets(ctx context.Context, input ListPolicySetsInput) (ListPolicySetsOutput, error) {
	apires, err := c.raw.ListPolicySets(ctx, connect.NewRequest(&authzv1alpha1.ListPolicySetsRequest{
		Universe:    "default",
		Environment: "production",
		PageToken:   input.PageToken,
	}))
	if err != nil {
		return ListPolicySetsOutput{}, err
	}

	res := ListPolicySetsOutput{
		NextPageToken: apires.Msg.NextPageToken,
		PolicySets:    make([]PolicySet, len(apires.Msg.PolicySets)),
	}

	for i, p := range apires.Msg.PolicySets {
		res.PolicySets[i] = PolicySetFromAPI(p)
	}

	return res, nil
}

type listPolicySetsRequestCall struct {
	input  ListPolicySetsInput
	client *Client
}

func (c *Client) ListPolicySetsRequest(input ListPolicySetsInput) *listPolicySetsRequestCall {
	return &listPolicySetsRequestCall{
		input:  input,
		client: c,
	}
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *listPolicySetsRequestCall) Pages(ctx context.Context, f func(ListPolicySetsOutput) error) error {
	// resets the input back to its original state
	originalPageToken := c.input.PageToken
	defer func() { c.input.PageToken = originalPageToken }()
	for {
		x, err := c.client.ListPolicySets(ctx, c.input)
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
