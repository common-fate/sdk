package provisioner

import (
	"github.com/common-fate/sdk/eid"
	"github.com/common-fate/sdk/service/entity"
)

type Grant struct {
	Grant           eid.EID  `json:"grant"`
	Principal       eid.EID  `json:"principal"`
	MappedPrincipal *eid.EID `json:"mapped_principal,omitempty"`
	Target          eid.EID  `json:"target"`
	Role            eid.EID  `json:"role"`
}

type GrantResponse struct {
	Entities entity.EntityJSON `json:"entities"`
}
