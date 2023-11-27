package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type ListPoliciesInput struct {
	PageToken string
}

type ListPolicyResult struct {
	Policies []Policy
	NextPage string
}

func (c *Client) ListPolicies(ctx context.Context, input ListPoliciesInput) (ListPolicyResult, error) {
	apires, err := c.raw.ListPolicies(ctx, connect.NewRequest(&authzv1alpha1.ListPoliciesRequest{
		Universe:    "default",
		Environment: "production",
		PageToken:   input.PageToken,
	}))
	if err != nil {
		return ListPolicyResult{}, err
	}

	res := ListPolicyResult{
		NextPage: apires.Msg.NextPageToken,
		Policies: make([]Policy, len(apires.Msg.Policies)),
	}

	for i, p := range apires.Msg.Policies {
		res.Policies[i] = Policy{
			ID:    p.Id,
			Cedar: p.Cedar,
		}
	}

	return res, nil
}
