package authz

import (
	"fmt"

	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

// UID is a unique identifier for an entity
// which includes both it's type and ID.
//
// e.g. 'GCP::Project::dev'
type UID struct {
	// Type of the entity.
	// e.g. in GCP::Project::dev
	// it is 'GCP::Project'
	Type string `json:"type"`

	// ID of the entity.
	// e.g. in GCP::Project::dev
	// it is 'dev'
	ID string `json:"id"`
}

func (u UID) ToAPI() *authzv1alpha1.UID {
	return &authzv1alpha1.UID{
		Type: u.Type,
		Id:   u.ID,
	}
}

func (u UID) String() string {
	return fmt.Sprintf(`%s::"%s"`, u.Type, u.ID)
}

func UIDFromAPI(uid *authzv1alpha1.UID) UID {
	return UID{
		Type: uid.Type,
		ID:   uid.Id,
	}
}
