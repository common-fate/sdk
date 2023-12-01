package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type PutEntityInput struct {
	Entity Entity
}

// PutEntity just currently calls the BatchPutEntity API with a single entity.
func (c *Client) PutEntity(ctx context.Context, input PutEntityInput) (*authzv1alpha1.BatchPutEntityResponse, error) {
	var req = &authzv1alpha1.BatchPutEntityRequest{
		Universe: "default",
		Entities: []*authzv1alpha1.Entity{},
	}

	parsed, children, err := transformToEntity(input.Entity)
	if err != nil {
		return nil, err
	}

	req.Entities = append(req.Entities, parsed)
	req.Children = append(req.Children, children...)

	res, err := c.raw.BatchPutEntity(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
