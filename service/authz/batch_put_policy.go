package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type BatchPutPolicySetInput struct {
	PolicySets []PolicySetInput
}

type BatchPutPolicySetOutput = authzv1alpha1.BatchPutPolicySetResponse

func (c *Client) BatchPutPolicySet(ctx context.Context, input BatchPutPolicySetInput) (*BatchPutPolicySetOutput, error) {
	p := make([]*authzv1alpha1.PolicySetInput, len(input.PolicySets))

	for i, ps := range input.PolicySets {
		p[i] = ps.ToAPI()
	}

	res, err := c.raw.BatchPutPolicySet(ctx, connect.NewRequest(&authzv1alpha1.BatchPutPolicySetRequest{
		Universe:    "default",
		Environment: "production",
		PolicySets:  p,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
