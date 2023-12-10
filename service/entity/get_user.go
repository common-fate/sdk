package entity

import (
	"context"

	"github.com/common-fate/sdk/eid"
	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
)

// GetUser hydrates a resource with the 'name' and 'email' attribute from the cache.
// This may result in additional API calls to the authz service if the entity is not
// present in the client cache.
func (c *Client) GetUser(ctx context.Context, eid eid.EID) *accessv1alpha1.User {
	return &accessv1alpha1.User{
		Eid:   eid.ToAPI(),
		Name:  c.AttrCache().String(ctx, eid, "name"),
		Email: c.AttrCache().String(ctx, eid, "email"),
	}
}
