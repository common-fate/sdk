package uiderr

import (
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/hashicorp/go-multierror"
)

// ErrorMap allows errors to be tracked when batch processing
// multiple entities at once.
type ErrorMap struct {
	Errors map[uid.UID]*multierror.Error
}

func Map() ErrorMap {
	return ErrorMap{Errors: map[uid.UID]*multierror.Error{}}
}

func (m *ErrorMap) Add(id uid.UID, err error) {
	existing, ok := m.Errors[id]
	if !ok {
		m.Errors[id] = multierror.Append(err)
	} else {
		m.Errors[id] = multierror.Append(existing, err)
	}
}
