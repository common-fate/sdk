package policyset

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

// Policy is an authorization policy used to make access decisions.
type Policy struct {
	// ID of the policy.
	ID string `json:"id"`
	// Text is the Cedar source code for the policy.
	Text string `json:"text"`
}

func (p *Policy) ToAPI() *authzv1alpha1.Policy {
	return &authzv1alpha1.Policy{Id: p.ID, Text: p.Text}
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

	Text string `json:"text"`
}

func FromAPI(input *authzv1alpha1.PolicySet) PolicySet {
	if input == nil {
		return PolicySet{}
	}

	ps := PolicySet{
		ID:   input.Id,
		Text: input.Text,
	}

	for _, p := range input.Policies {
		ps.Policies = append(ps.Policies, PolicyFromAPI(p))
	}

	return ps
}

func (ps *PolicySet) ToAPI() *authzv1alpha1.PolicySet {
	out := &authzv1alpha1.PolicySet{
		Id:       ps.ID,
		Policies: make([]*authzv1alpha1.Policy, len(ps.Policies)),
		Text:     ps.Text,
	}
	for i, p := range ps.Policies {
		out.Policies[i] = p.ToAPI()
	}
	return out
}
