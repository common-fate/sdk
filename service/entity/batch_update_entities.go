package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
)

type BatchUpdateInput struct {
	Entities []Entity
	Children []ChildRelation
}

func (c *Client) BatchUpdateEntities(ctx context.Context, input BatchUpdateInput) (*entityv1alpha1.BatchUpdateResponse, error) {
	var req = &entityv1alpha1.BatchUpdateRequest{
		Universe:    "default",
		PutEntities: []*entityv1alpha1.Entity{},
	}

	for _, e := range input.Entities {
		parsed, children, err := EntityToAPI(e)
		if err != nil {
			return nil, err
		}

		req.PutEntities = append(req.PutEntities, parsed)
		req.PutChildren = append(req.PutChildren, children...)
	}

	for _, c := range input.Children {
		req.PutChildren = append(req.PutChildren, c.ToAPI())
	}

	res, err := c.raw.BatchUpdate(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
