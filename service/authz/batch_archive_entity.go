package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
)

type BatchArchiveEntityInput struct {
	Archive   []uid.UID
	Unarchive []uid.UID
}

// BatchArchiveEntity soft-deletes entities so that by default they will no longer be present in the ListEntities requests.
func (c *Client) BatchArchiveEntity(ctx context.Context, input BatchArchiveEntityInput) (*authzv1alpha1.BatchArchiveEntityResponse, error) {
	apiArchiveUIDs := make([]*authzv1alpha1.UID, len(input.Archive))

	for i, uid := range input.Archive {
		apiArchiveUIDs[i] = uid.ToAPI()
	}

	apiUnarchiveUIDs := make([]*authzv1alpha1.UID, len(input.Unarchive))

	for i, uid := range input.Unarchive {
		apiUnarchiveUIDs[i] = uid.ToAPI()
	}

	res, err := c.raw.BatchArchiveEntity(ctx, connect.NewRequest(&authzv1alpha1.BatchArchiveEntityRequest{
		Universe:  "default",
		Archive:   apiArchiveUIDs,
		Unarchive: apiUnarchiveUIDs,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
