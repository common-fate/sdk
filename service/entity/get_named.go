package entity

import (
	"context"

	"github.com/common-fate/sdk/eid"
	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
)

// GetNamed hydrates a resource with the 'name' attribute from the cache.
// This may result in additional API calls to the authz service if the entity is not
// present in the client cache.
func (c *Client) GetNamed(ctx context.Context, eid eid.EID) *accessv1alpha1.NamedEID {
	return &accessv1alpha1.NamedEID{
		Eid:  eid.ToAPI(),
		Name: c.AttrCache().String(ctx, eid, "name"),
	}
}
