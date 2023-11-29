package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/patrickmn/go-cache"
)

type GetEntityInput struct {
	UID uid.UID
}

func (c *Client) GetEntity(ctx context.Context, input GetEntityInput) (*authzv1alpha1.GetEntityResponse, error) {
	req := &authzv1alpha1.GetEntityRequest{
		Universe: "default",
		Entity:   input.UID.ToAPI(),
	}

	res, err := c.raw.GetEntity(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	c.cache.Set(input.UID.String(), res.Msg.Entity, cache.DefaultExpiration)

	return res.Msg, nil
}
