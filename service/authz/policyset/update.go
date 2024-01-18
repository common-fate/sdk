package policyset

import (
	"context"

	"connectrpc.com/connect"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type UpdateInput struct {
	PolicySet Input
}

type UpdateOutput = authzv1alpha1.UpdatePolicySetResponse

func (c *Client) Update(ctx context.Context, input UpdateInput) (*UpdateOutput, error) {
	res, err := c.raw.UpdatePolicySet(ctx, connect.NewRequest(&authzv1alpha1.UpdatePolicySetRequest{
		PolicySet: input.PolicySet.ToAPI(),
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
