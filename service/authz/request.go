package authz

import authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"

type Request struct {
	Principal UID `json:"principal"`
	Action    UID `json:"action"`
	Resource  UID `json:"resource"`
}

func (r Request) ToAPI(key string) *authzv1alpha1.Request {
	return &authzv1alpha1.Request{
		ClientKey: key,
		Principal: r.Principal.ToAPI(),
		Action:    r.Action.ToAPI(),
		Resource:  r.Resource.ToAPI(),
	}
}
