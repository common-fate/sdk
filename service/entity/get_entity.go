package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/patrickmn/go-cache"
)

type GetInput struct {
	UID uid.UID
	// UseCache will try and retrieve entities from
	// the client cache if it's present.
	//
	// If all entities are cached, no API call will be made.
	UseCache bool
}

func (c *Client) GetEntity(ctx context.Context, input GetInput) (*entityv1alpha1.GetResponse, error) {
	req := &entityv1alpha1.GetRequest{
		Universe: "default",
		Uid:      input.UID.ToAPI(),
	}

	if input.UseCache {
		if got, ok := c.cache.Get(input.UID.String()); ok {
			if entity, ok := got.(*entityv1alpha1.Entity); ok {
				return &entityv1alpha1.GetResponse{
					Entity: entity,
				}, nil
			}
		}
	}

	res, err := c.raw.Get(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	c.cache.Set(input.UID.String(), res.Msg.Entity, cache.DefaultExpiration)

	return res.Msg, nil
}
