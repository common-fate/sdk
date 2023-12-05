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
	PolicySets []PolicySet
	NextPage   string
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
		NextPage:   apires.Msg.NextPageToken,
		PolicySets: make([]PolicySet, len(apires.Msg.PolicySets)),
	}

	for i, p := range apires.Msg.PolicySets {
		res.PolicySets[i] = PolicySetFromAPI(p)
	}

	return res, nil
}
