package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

// removes entities from the authorization service.
func (c *Client) BatchDeleteEntity(ctx context.Context, uids ...UID) (*authzv1alpha1.BatchDeleteEntityResponse, error) {
	apiUIDs := make([]*authzv1alpha1.UID, len(uids))

	for i, uid := range uids {
		apiUIDs[i] = uid.ToAPI()
	}

	res, err := c.raw.BatchDeleteEntity(ctx, connect.NewRequest(&authzv1alpha1.BatchDeleteEntityRequest{
		Universe: "default",
		Entities: apiUIDs,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
