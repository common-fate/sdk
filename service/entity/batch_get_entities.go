package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/patrickmn/go-cache"
)

type BatchGetInput struct {
	EIDs []eid.EID
	// UseCache will try and retrieve entities from
	// the client cache if it's present.
	//
	// If all entities are cached, no API call will be made.
	UseCache bool
}

type BatchGetOutput = entityv1alpha1.BatchGetResponse

func (c *Client) BatchGet(ctx context.Context, input BatchGetInput) (*BatchGetOutput, error) {
	// de-duplicate EIDs
	eids := map[eid.EID]bool{}

	for _, i := range input.EIDs {
		eids[i] = true
	}

	req := &entityv1alpha1.BatchGetRequest{
		Universe: "default",
	}

	var cached []*entityv1alpha1.Entity

	for u := range eids {
		if input.UseCache {
			if got, ok := c.cache.Get(u.String()); ok {
				if entity, ok := got.(*entityv1alpha1.Entity); ok {
					cached = append(cached, entity)
					continue
				}
			}
		}

		req.Eids = append(req.Eids, u.ToAPI())
	}

	output := &entityv1alpha1.BatchGetResponse{}

	if len(req.Eids) > 0 {
		res, err := c.raw.BatchGet(ctx, connect.NewRequest(req))
		if err != nil {
			return nil, err
		}
		output = res.Msg

		// update the attribute cache
		for _, e := range res.Msg.Entities {
			eid := eid.EID{
				Type: e.Eid.Type,
				ID:   e.Eid.Id,
			}

			c.cache.Set(eid.String(), e, cache.DefaultExpiration)
		}
	}

	// add any cached entities we retrieved
	output.Entities = append(output.Entities, cached...)

	return output, nil
}
