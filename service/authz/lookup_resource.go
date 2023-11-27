package authz

type LookupResourcesInput struct {
	Principal    UID
	ResourceType string
	PageToken    string
}

// func (c *Client) LookupResources(ctx context.Context, input LookupResourcesInput) (*authzv1alpha1.LookupResourcesResponse, error) {
// 	res, err := c.raw.LookupResources(ctx, connect.NewRequest(&authzv1alpha1.LookupResourcesRequest{
// 		Universe:     "default",
// 		Environment:  "production",
// 		Principal:    input.Principal.ToAPI(),
// 		ResourceType: input.ResourceType,
// 		PageToken:    input.PageToken,
// 	}))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res.Msg, nil
// }
