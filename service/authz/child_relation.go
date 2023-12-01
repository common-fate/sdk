package authz

import (
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
)

type ChildRelation struct {
	Parent uid.UID
	Child  uid.UID
}

func (cr ChildRelation) ToAPI() *authzv1alpha1.ChildRelation {
	return &authzv1alpha1.ChildRelation{
		Parent: cr.Parent.ToAPI(),
		Child:  cr.Child.ToAPI(),
	}
}
