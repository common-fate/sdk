package authz

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

// PolicySetInput is the input to create a PolicySet
type PolicySetInput struct {
	// ID of the policy.
	ID string `json:"id"`
	// Text is the Cedar source code for the policy.
	Text string `json:"text"`
}

func (p PolicySetInput) ToAPI() *authzv1alpha1.PolicySetInput {
	return &authzv1alpha1.PolicySetInput{
		Id:   p.ID,
		Text: p.Text,
	}
}
