package policyset

import "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1/authzv1alpha1connect"

type Client struct {
	raw authzv1alpha1connect.PolicyServiceClient
}
