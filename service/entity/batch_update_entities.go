package entity

import (
	"context"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
)

type BatchUpdateInput struct {
	// Put is a slice of entities to be upserted to the entity store.
	Put []Entity
	// PutChildren is a slice of parent/child relationships to be added to the entity store.
	PutChildren []ChildRelation
	// Archive is a slice of UIDs of entities to be archived.
	// When an entity is archived it will no longer appear in the List API calls by default.
	Archive []uid.UID
	// Unarchive is a slice of UIDs of entities to be made active again.
	Unarchive []uid.UID
	// Unarchive is a slice of UIDs of entities to be deleted from the entity store.
	Delete []uid.UID
	// DeleteChildren is a slice of parent/child relations to be removed from the entity store.
	DeleteChildren []ChildRelation
}

type BatchUpdateOutput = entityv1alpha1.BatchUpdateResponse

func (c *Client) BatchUpdate(ctx context.Context, input BatchUpdateInput) (*BatchUpdateOutput, error) {
	var req = &entityv1alpha1.BatchUpdateRequest{
		Universe:    "default",
		PutEntities: []*entityv1alpha1.Entity{},
	}

	for _, e := range input.Put {
		parsed, children, err := EntityToAPI(e)
		if err != nil {
			return nil, err
		}

		req.PutEntities = append(req.PutEntities, parsed)
		req.PutChildren = append(req.PutChildren, children...)
	}

	for _, c := range input.PutChildren {
		req.PutChildren = append(req.PutChildren, c.ToAPI())
	}

	for _, c := range input.Archive {
		req.ArchiveEntities = append(req.ArchiveEntities, c.ToAPI())
	}

	for _, c := range input.Unarchive {
		req.UnarchiveEntities = append(req.UnarchiveEntities, c.ToAPI())
	}

	for _, c := range input.Delete {
		req.DeleteEntities = append(req.DeleteEntities, c.ToAPI())
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
