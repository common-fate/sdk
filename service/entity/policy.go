package entity

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

// Policy is an authorization policy used to make access decisions.
type Policy struct {
	// ID of the policy.
	ID string `json:"id"`
	// Cedar code for the policy.
	Cedar string `json:"cedar"`
}

func (p Policy) ToAPI() *authzv1alpha1.Policy {
	return &authzv1alpha1.Policy{
		Id:    p.ID,
		Cedar: p.Cedar,
	}
}
