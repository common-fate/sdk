package batchauthz

import (
	"context"
	"fmt"
	"testing"

	"github.com/common-fate/sdk/service/authz"
	"github.com/common-fate/sdk/service/authz/uid"
)

// MockBatch is a mock batch authorizer which implements the
// same methods, used for testing
type MockBatch struct {
	t           *testing.T
	evaluations map[uid.UID]map[uid.UID]map[uid.UID]Evaluation
	executed    bool
	strict      bool
}

var _ Authorizer = &MockBatch{}

func NewMock(t *testing.T) *MockBatch {
	return &MockBatch{
		evaluations: map[uid.UID]map[uid.UID]map[uid.UID]Evaluation{},
	}
}

// Strict enables strict evaluation for the mock authorizer.
// In strict evaluation mode, any requests that have not been mocked
// will cause the test to fail
func (a *MockBatch) Strict() *MockBatch {
	a.strict = true
	return a
}

func (a *MockBatch) Allow(req authz.Request) *MockBatch {
	a.Mock(req, Evaluation{Allowed: true})
	return a
}
func (a *MockBatch) AllowActions(principal uid.UID, resource uid.UID, actions ...uid.UID) *MockBatch {
	for _, action := range actions {
		req := authz.Request{
			Principal: principal,
			Action:    action,
			Resource:  resource,
		}
		a.Mock(req, Evaluation{Allowed: true})
	}

	return a
}

// Mock a particular request to return the specified evaluation
func (a *MockBatch) Mock(req authz.Request, eval Evaluation) *MockBatch {
	principalMap, ok := a.evaluations[req.Principal]
	if !ok {
		principalMap = make(map[uid.UID]map[uid.UID]Evaluation)
		a.evaluations[req.Principal] = principalMap
	}

	// Check if the resource map exists for the principal
	resourceMap, ok := principalMap[req.Resource]
	if !ok {
		resourceMap = make(map[uid.UID]Evaluation)
		principalMap[req.Resource] = resourceMap
	}

	_, ok = resourceMap[req.Action]
	if ok {
		// there were multiple evaluations for a particular principal/action/resource combination.
		// this is a logic error and so we return an error
		a.t.Fatalf("multiple authorization evaluations found for principal %s, action %s and resource %s: ensure that Mock is only called once for each combination of principal/action/resource", req.Principal, req.Action, req.Resource)
	}

	// Populate the evaluations map
	resourceMap[req.Action] = eval

	return a
}

func (a *MockBatch) AddRequest(req authz.Request) Authorizer {
	return a
}

func (a *MockBatch) Authorize(ctx context.Context) error {
	a.executed = true
	return nil
}

// AuthorizeWasCalled returns true if Authorize() was called
func (a *MockBatch) AuthorizeWasCalled() bool {
	return a.executed
}

func (a *MockBatch) IsPermitted(req authz.Request) (Evaluation, error) {
	// Check if the principal exists in the evaluations map
	principalMap, principalExists := a.evaluations[req.Principal]
	if !principalExists {
		if !a.strict {
			return Evaluation{Allowed: false}, nil
		}

		return Evaluation{}, fmt.Errorf("principal %s not found in evaluations", req.Principal)
	}

	// Check if the resource map exists for the principal
	resourceMap, resourceExists := principalMap[req.Resource]
	if !resourceExists {
		if !a.strict {
			return Evaluation{Allowed: false}, nil
		}

		return Evaluation{}, fmt.Errorf("resource %s not found for the principal %s", req.Resource, req.Principal)
	}

	// Get the evaluation for the specified action
	eval, evalExists := resourceMap[req.Action]
	if !evalExists {
		if !a.strict {
			return Evaluation{Allowed: false}, nil
		}

		return Evaluation{}, fmt.Errorf("evaluation not found for the specified action %s on resource %s for principal %s", req.Action, req.Resource, req.Principal)
	}

	return eval, nil
}
