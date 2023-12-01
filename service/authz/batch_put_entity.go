package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type BatchPutEntityInput struct {
	Entities []Entity
	Children []ChildRelation
}

func (c *Client) BatchPutEntity(ctx context.Context, input BatchPutEntityInput) (*authzv1alpha1.BatchPutEntityResponse, error) {
	var req = &authzv1alpha1.BatchPutEntityRequest{
		Universe: "default",
		Entities: []*authzv1alpha1.Entity{},
	}

	for _, e := range input.Entities {
		parsed, children, err := transformToEntity(e)
		if err != nil {
			return nil, err
		}

		req.Entities = append(req.Entities, parsed)
		req.Children = append(req.Children, children...)
	}

	for _, c := range input.Children {
		req.Children = append(req.Children, c.ToAPI())
	}

	res, err := c.raw.BatchPutEntity(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
