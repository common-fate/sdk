package entity

import (
	"github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1/entityv1alpha1connect"
)

// RawClient returns the underlying Buf Connect client.
func (c *Client) RawClient() entityv1alpha1connect.EntityServiceClient {
	return c.raw
}
