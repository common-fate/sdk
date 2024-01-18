package entity

import (
	"context"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
)

type BatchUpdateInput struct {
	// Put is a slice of entities to be upserted to the entity store.
	Put []Entity
	// PutChildren is a slice of parent/child relationships to be added to the entity store.
	PutChildren []ChildRelation
	// Archive is a slice of EIDs of entities to be archived.
	// When an entity is archived it will no longer appear in the List API calls by default.
	Archive []eid.EID
	// Unarchive is a slice of EIDs of entities to be made active again.
	Unarchive []eid.EID
	// Unarchive is a slice of EIDs of entities to be deleted from the entity store.
	Delete []eid.EID
	// DeleteChildren is a slice of parent/child relations to be removed from the entity store.
	DeleteChildren []ChildRelation
}

type BatchUpdateOutput = entityv1alpha1.BatchUpdateResponse

func (c *Client) BatchUpdate(ctx context.Context, input BatchUpdateInput) (*BatchUpdateOutput, error) {
	var req = &entityv1alpha1.BatchUpdateRequest{
		Universe: "default",
		Put:      []*entityv1alpha1.Entity{},
	}

	if len(input.Put) == 0 && len(input.PutChildren) == 0 && len(input.Archive) == 0 && len(input.Unarchive) == 0 && len(input.Delete) == 0 && len(input.DeleteChildren) == 0 {
		// no API call required
		return nil, nil
	}

	for _, e := range input.Put {
		parsed, children, err := Marshal(e)
		if err != nil {
			return nil, err
		}

		req.Put = append(req.Put, parsed)
		req.PutChildren = append(req.PutChildren, children...)
	}

	for _, c := range input.PutChildren {
		req.PutChildren = append(req.PutChildren, c.ToAPI())
	}

	for _, c := range input.Archive {
		req.Archive = append(req.Archive, c.ToAPI())
	}

	for _, c := range input.Unarchive {
		req.Unarchive = append(req.Unarchive, c.ToAPI())
	}

	for _, c := range input.Delete {
		req.Delete = append(req.Delete, c.ToAPI())
	}

	for _, c := range input.DeleteChildren {
		req.DeleteChildren = append(req.DeleteChildren, c.ToAPI())
	}

	res, err := c.raw.BatchUpdate(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
