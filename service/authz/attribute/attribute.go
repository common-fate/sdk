package attribute

import (
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz"
)

func String(key string, value string) authz.Attribute {
	return strAttr{Key: key, Value: value}
}

func Bool(key string, value bool) authz.Attribute {
	return boolAttr{Key: key, Value: value}
}

func Int(key string, value int) authz.Attribute {
	return intAttr{Key: key, Value: value}
}

type strAttr struct {
	Key   string
	Value string
}

func (s strAttr) ToAPI() *authzv1alpha1.Attribute {
	return &authzv1alpha1.Attribute{
		Key: s.Key,
		Value: &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Str{
				Str: s.Value,
			},
		},
	}
}

type boolAttr struct {
	Key   string
	Value bool
}

func (a boolAttr) ToAPI() *authzv1alpha1.Attribute {
	return &authzv1alpha1.Attribute{
		Key: a.Key,
		Value: &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Bool{
				Bool: a.Value,
			},
		},
	}
}

type intAttr struct {
	Key   string
	Value int
}

func (a intAttr) ToAPI() *authzv1alpha1.Attribute {
	return &authzv1alpha1.Attribute{
		Key: a.Key,
		Value: &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Long{
				Long: int64(a.Value),
			},
		},
	}
}
