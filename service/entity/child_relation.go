package entity

import (
	"github.com/common-fate/sdk/eid"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
)

type ChildRelation struct {
	Parent eid.EID
	Child  eid.EID
}

func (cr ChildRelation) ToAPI() *entityv1alpha1.ChildRelation {
	return &entityv1alpha1.ChildRelation{
		Parent: cr.Parent.ToAPI(),
		Child:  cr.Child.ToAPI(),
	}
}
