package policyset

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type CreateInput struct {
	PolicySet Input
}

type CreateOutput = authzv1alpha1.CreatePolicySetResponse

func (c *Client) Create(ctx context.Context, input CreateInput) (*CreateOutput, error) {
	res, err := c.raw.CreatePolicySet(ctx, connect.NewRequest(&authzv1alpha1.CreatePolicySetRequest{
		PolicySet: input.PolicySet.ToAPI(),
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
