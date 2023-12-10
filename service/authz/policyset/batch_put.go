package policyset

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type BatchPutInput struct {
	PolicySets []Input
}

type BatchPutOutput = authzv1alpha1.BatchPutPolicySetResponse

func (c *Client) BatchPut(ctx context.Context, input BatchPutInput) (*BatchPutOutput, error) {
	p := make([]*authzv1alpha1.PolicySetInput, len(input.PolicySets))

	for i, ps := range input.PolicySets {
		p[i] = ps.ToAPI()
	}

	res, err := c.raw.BatchPutPolicySet(ctx, connect.NewRequest(&authzv1alpha1.BatchPutPolicySetRequest{
		PolicySets: p,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
