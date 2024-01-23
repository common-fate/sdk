package policyset

import (
	"context"

	"connectrpc.com/connect"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type GetInput struct {
	ID string
}

type GetOutput struct {
	PolicySet PolicySet `json:"policySet"`
}

func (c *Client) Get(ctx context.Context, input GetInput) (GetOutput, error) {
	apires, err := c.raw.GetPolicySet(ctx, connect.NewRequest(&authzv1alpha1.GetPolicySetRequest{
		Id: input.ID,
	}))
	if err != nil {
		return GetOutput{}, err
	}

	res := GetOutput{
		PolicySet: FromAPI(apires.Msg.PolicySet),
	}

	return res, nil
}
