package entity

import (
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
)

type ChildRelation struct {
	Parent uid.UID
	Child  uid.UID
}

func (cr ChildRelation) ToAPI() *entityv1alpha1.ChildRelation {
	return &entityv1alpha1.ChildRelation{
		Parent: cr.Parent.ToAPI(),
		Child:  cr.Child.ToAPI(),
	}
}
