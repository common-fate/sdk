package provisioner

import (
	"github.com/common-fate/sdk/eid"
	"github.com/common-fate/sdk/service/entity"
)

type Context struct {
	Attr1 string `json:"attr_1"`
	Attr2 string `json:"attr_2"`
	Attr3 string `json:"attr_3"`
	Attr4 string `json:"attr_4"`
	Attr5 string `json:"attr_5"`
}

type Grant struct {
	Grant              eid.EID  `json:"grant"`
	Principal          eid.EID  `json:"principal"`
	MappedPrincipal    *eid.EID `json:"mapped_principal,omitempty"`
	Target             eid.EID  `json:"target"`
	Role               eid.EID  `json:"role"`
	ProvisionerContext Context  `json:"provisioner_context"`
}

type GrantResponse struct {
	Entities entity.EntityJSON `json:"entities"`
}
