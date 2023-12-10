package policyset

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

// Input is the input to create a PolicySet
type Input struct {
	// ID of the policy.
	ID string `json:"id"`
	// Text is the Cedar source code for the policy.
	Text string `json:"text"`
}

func (p Input) ToAPI() *authzv1alpha1.PolicySetInput {
	return &authzv1alpha1.PolicySetInput{
		Id:   p.ID,
		Text: p.Text,
	}
}
