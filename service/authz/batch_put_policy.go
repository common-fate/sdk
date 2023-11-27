package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type BatchPutPolicyInput struct {
	Policies []Policy
}

func (c *Client) BatchPutPolicy(ctx context.Context, input BatchPutPolicyInput) (*authzv1alpha1.BatchPutPolicyResponse, error) {
	p := make([]*authzv1alpha1.Policy, len(input.Policies))

	for i, uid := range input.Policies {
		p[i] = uid.ToAPI()
	}

	res, err := c.raw.BatchPutPolicy(ctx, connect.NewRequest(&authzv1alpha1.BatchPutPolicyRequest{
		Universe:    "default",
		Environment: "production",
		Policies:    p,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
