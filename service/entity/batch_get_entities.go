package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/patrickmn/go-cache"
)

type BatchGetInput struct {
	UIDs []uid.UID
	// UseCache will try and retrieve entities from
	// the client cache if it's present.
	//
	// If all entities are cached, no API call will be made.
	UseCache bool
}

type BatchGetOutput = entityv1alpha1.BatchGetResponse

func (c *Client) BatchGet(ctx context.Context, input BatchGetInput) (*BatchGetOutput, error) {
	req := &entityv1alpha1.BatchGetRequest{
		Universe: "default",
	}

	var cached []*entityv1alpha1.Entity

	for _, u := range input.UIDs {
		if input.UseCache {
			if got, ok := c.cache.Get(u.String()); ok {
				if entity, ok := got.(*entityv1alpha1.Entity); ok {
					cached = append(cached, entity)
					continue
				}
			}
		}

		req.Uids = append(req.Uids, u.ToAPI())
	}

	output := &entityv1alpha1.BatchGetResponse{}

	if len(req.Uids) > 0 {
		res, err := c.raw.BatchGet(ctx, connect.NewRequest(req))
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
