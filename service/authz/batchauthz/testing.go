package batchauthz

import (
	"context"
	"fmt"
	"testing"

	"github.com/common-fate/sdk/eid"
	"github.com/common-fate/sdk/service/authz"
)

// MockBatch is a mock batch authorizer which implements the
// same methods, used for testing
type MockBatch struct {
	t           *testing.T
	lastRequest *authz.Request
	evaluations map[eid.EID]map[eid.EID]map[eid.EID]Evaluation
	executed    bool
	strict      bool
}

var _ Authorizer = &MockBatch{}

func NewMock(t *testing.T) *MockBatch {
	return &MockBatch{
		evaluations: map[eid.EID]map[eid.EID]map[eid.EID]Evaluation{},
	}
}

// Strict enables strict evaluation for the mock authorizer.
// In strict evaluation mode, any requests that have not been mocked
// will cause the test to fail
func (a *MockBatch) Strict() *MockBatch {
	a.strict = true
	return a
}

func (a *MockBatch) Allow(principal eid.EID, resource eid.EID, actions ...eid.EID) *MockBatch {
	for _, action := range actions {
		req := authz.Request{
			Principal: principal,
			Action:    action,
			Resource:  resource,
		}
		a.lastRequest = &req
		a.Mock(req, Evaluation{Allowed: true})
	}

	return a
}

// Annotations adds annotations to the last Allow() call. Fails the test if called without first calling Allow().
func (a *MockBatch) Annotations(key string, values ...string) *MockBatch {
	if a.lastRequest == nil {
		a.t.Fatal("annotations must be called after Allow(), for example: batchauthz.NewMock().Allow(principal, resource, action).Annotations()")
	}

	eval, err := a.getMockEvaluation(*a.lastRequest)
	if err != nil {
		a.t.Fatalf("mockbatch: error adding annotations: %s", err)
	}

	for _, v := range values {
		eval.Annotations.Set(key, Annotation{Value: v})
	}

	return a.Mock(*a.lastRequest, eval)
}

// Mock a particular request to return the specified evaluation
func (a *MockBatch) Mock(req authz.Request, eval Evaluation) *MockBatch {
	principalMap, ok := a.evaluations[req.Principal]
	if !ok {
		principalMap = make(map[eid.EID]map[eid.EID]Evaluation)
		a.evaluations[req.Principal] = principalMap
	}

	// Check if the resource map exists for the principal
	resourceMap, ok := principalMap[req.Resource]
	if !ok {
		resourceMap = make(map[eid.EID]Evaluation)
		principalMap[req.Resource] = resourceMap
	}

	// Populate the evaluations map
	resourceMap[req.Action] = eval

	return a
}

func (a *MockBatch) AddRequest(req authz.Request) error {
	return nil
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
	return a.getMockEvaluation(req)
}

func (a *MockBatch) getMockEvaluation(req authz.Request) (Evaluation, error) {
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
