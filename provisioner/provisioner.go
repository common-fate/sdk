package provisioner

import (
	"github.com/common-fate/sdk/service/entity"
	"github.com/common-fate/sdk/uid"
)

type Grant struct {
	Grant           uid.UID  `json:"grant"`
	Principal       uid.UID  `json:"principal"`
	MappedPrincipal *uid.UID `json:"mapped_principal,omitempty"`
	Target          uid.UID  `json:"target"`
	Role            uid.UID  `json:"role"`
}

type GrantResponse struct {
	Entities entity.EntityJSON `json:"entities"`
}
