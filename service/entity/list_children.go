package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
)

type ListChildrenInput struct {
	Parent    uid.UID
	Type      string
	PageToken string
}

type ListChildrenOutput = entityv1alpha1.ListChildrenResponse

func (c *Client) ListChildren(ctx context.Context, input ListChildrenInput) (*ListChildrenOutput, error) {
	req := &entityv1alpha1.ListChildrenRequest{
		Universe:  "default",
		PageToken: input.PageToken,
		Parent:    input.Parent.ToAPI(),
		Type:      input.Type,
	}

	res, err := c.raw.ListChildren(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
