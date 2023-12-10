package eiderr

import (
	"github.com/common-fate/sdk/eid"
	"github.com/hashicorp/go-multierror"
)

// ErrorMap allows errors to be tracked when batch processing
// multiple entities at once.
type ErrorMap struct {
	Errors map[eid.EID]*multierror.Error
}

func Map() ErrorMap {
	return ErrorMap{Errors: map[eid.EID]*multierror.Error{}}
}

func (m *ErrorMap) Add(id eid.EID, err error) {
	existing, ok := m.Errors[id]
	if !ok {
		m.Errors[id] = multierror.Append(err)
	} else {
		m.Errors[id] = multierror.Append(existing, err)
	}
}
