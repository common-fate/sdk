package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/patrickmn/go-cache"
)

type BatchGetEntityInput struct {
	UIDs []uid.UID
	// UseCache will try and retrieve entities from
	// the client cache if it's present.
	//
	// If all entities are cached, no API call will be made.
	UseCache bool
}

func (c *Client) BatchGetEntity(ctx context.Context, input BatchGetEntityInput) (*authzv1alpha1.BatchGetEntityResponse, error) {
	req := &authzv1alpha1.BatchGetEntityRequest{
		Universe: "default",
	}

	var cached []*authzv1alpha1.Entity

	for _, u := range input.UIDs {
		if input.UseCache {
			if got, ok := c.cache.Get(u.String()); ok {
				if entity, ok := got.(*authzv1alpha1.Entity); ok {
					cached = append(cached, entity)
					continue
				}
			}
		}

		req.Entities = append(req.Entities, u.ToAPI())
	}

	output := &authzv1alpha1.BatchGetEntityResponse{}

	if len(req.Entities) > 0 {
		res, err := c.raw.BatchGetEntity(ctx, connect.NewRequest(req))
		if err != nil {
			return nil, err
		}
		output = res.Msg

		// update the attribute cache
		for _, e := range res.Msg.Entities {
			uid := uid.UID{
				Type: e.Uid.Type,
				ID:   e.Uid.Id,
			}

			c.cache.Set(uid.String(), e, cache.DefaultExpiration)
		}
	}

	// add any cached entities we retrieved
	output.Entities = append(output.Entities, cached...)

	return output, nil
}
