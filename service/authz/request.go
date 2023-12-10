package authz

import (
	"github.com/common-fate/sdk/eid"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/entity"
)

type Request struct {
	Principal       eid.EID                `json:"principal"`
	Action          eid.EID                `json:"action"`
	Resource        eid.EID                `json:"resource"`
	OverlayEntities []entity.Entity        `json:"overlay_entities,omitempty"`
	OverlayChildren []entity.ChildRelation `json:"overlay_children,omitempty"`
}

func (r Request) ToAPI(key string) (*authzv1alpha1.Request, error) {
	res := &authzv1alpha1.Request{
		ClientKey: key,
		Principal: r.Principal.ToAPI(),
		Action:    r.Action.ToAPI(),
		Resource:  r.Resource.ToAPI(),
	}

	for _, e := range r.OverlayEntities {
		entity, children, err := entity.Marshal(e)
		if err != nil {
			return nil, err
		}
		res.OverlayEntities = append(res.OverlayEntities, entity)
		res.OverlayChildren = append(res.OverlayChildren, children...)
	}

	for _, c := range r.OverlayChildren {
		res.OverlayChildren = append(res.OverlayChildren, c.ToAPI())
	}

	return res, nil
}
