package policyset

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

// Policy is an authorization policy used to make access decisions.
type Policy struct {
	// ID of the policy.
	ID string `json:"id"`
	// Text is the Cedar source code for the policy.
	Text string `json:"text"`
}

func PolicyFromAPI(input *authzv1alpha1.Policy) Policy {
	if input == nil {
		return Policy{}
	}

	return Policy{
		ID:   input.Id,
		Text: input.Text,
	}
}

// PolicySet is a set of Policies
type PolicySet struct {
	// ID of the policy.
	ID string `json:"id"`

	Policies []Policy `json:"policies"`
}

func FromAPI(input *authzv1alpha1.PolicySet) PolicySet {
	if input == nil {
		return PolicySet{}
	}

	ps := PolicySet{
		ID: input.Id,
	}

	for _, p := range input.Policies {
		ps.Policies = append(ps.Policies, PolicyFromAPI(p))
	}

	return ps
}
