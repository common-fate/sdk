package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
)

type PutInput struct {
	Entity Entity
}

// PutEntity just currently calls the BatchUpdate API with a single entity.
func (c *Client) Put(ctx context.Context, input PutInput) (*BatchUpdateOutput, error) {
	var req = &entityv1alpha1.BatchUpdateRequest{
		Universe: "default",
	}

	parsed, children, err := EntityToAPI(input.Entity)
	if err != nil {
		return nil, err
	}

	req.Put = append(req.Put, parsed)
	req.PutChildren = append(req.PutChildren, children...)

	res, err := c.raw.BatchUpdate(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
