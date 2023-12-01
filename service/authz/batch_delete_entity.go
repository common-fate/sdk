package authz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
)

type BatchDeleteEntityInput struct {
	Entities []uid.UID
	Children []ChildRelation
}

// removes entities from the authorization service.
func (c *Client) BatchDeleteEntity(ctx context.Context, input BatchDeleteEntityInput) (*authzv1alpha1.BatchDeleteEntityResponse, error) {
	apiUIDs := make([]*authzv1alpha1.UID, len(input.Entities))

	for i, uid := range input.Entities {
		apiUIDs[i] = uid.ToAPI()
	}

	apiChildren := make([]*authzv1alpha1.ChildRelation, len(input.Children))

	for i, c := range input.Children {
		apiChildren[i] = c.ToAPI()
	}

	res, err := c.raw.BatchDeleteEntity(ctx, connect.NewRequest(&authzv1alpha1.BatchDeleteEntityRequest{
		Universe: "default",
		Entities: apiUIDs,
		Children: apiChildren,
	}))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}
