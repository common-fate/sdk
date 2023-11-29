package authz

import (
	"context"

	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
)

// GetNamed hydrates a resource with the 'name' attribute from the cache.
// This may result in additional API calls to the authz service if the entity is not
// present in the client cache.
func (c *Client) GetNamed(ctx context.Context, uid uid.UID) *accessv1alpha1.NamedUID {
	return &accessv1alpha1.NamedUID{
		Uid:  uid.ToAPI(),
		Name: c.AttrCache().String(ctx, uid, "name"),
	}
}
