package batchauthz

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/common-fate/sdk/service/authz"
)

// MockBatch is a mock batch authorizer which implements the
// same methods, used for testing
type MockBatch struct {
	t           *testing.T
	evaluations map[authz.UID]map[authz.UID]map[authz.UID]Evaluation
	executed    bool
}

func NewMock(t *testing.T) *MockBatch {
	return &MockBatch{
		evaluations: map[authz.UID]map[authz.UID]map[authz.UID]Evaluation{},
	}
}

// Mock a particular request to return the specified evaluation
func (a *MockBatch) Mock(req Request, eval Evaluation) *MockBatch {
	principalMap, ok := a.evaluations[req.Principal]
	if !ok {
		principalMap = make(map[authz.UID]map[authz.UID]Evaluation)
		a.evaluations[req.Principal] = principalMap
	}

	// Check if the resource map exists for the principal
	resourceMap, ok := principalMap[req.Resource]
	if !ok {
		resourceMap = make(map[authz.UID]Evaluation)
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

func (a *MockBatch) Execute(ctx context.Context, executor Executor) error {
	a.executed = true
	return nil
}

func (a *MockBatch) IsPermitted(req Request) (Evaluation, error) {
	if !a.executed {
		return Evaluation{}, errors.New("you must call Execute() to call the authz service before calling IsPermitted()")
	}

	// Check if the principal exists in the evaluations map
	principalMap, principalExists := a.evaluations[req.Principal]
	if !principalExists {
		return Evaluation{}, fmt.Errorf("principal %s not found in evaluations", req.Principal)
	}

	// Check if the resource map exists for the principal
	resourceMap, resourceExists := principalMap[req.Resource]
	if !resourceExists {
		return Evaluation{}, fmt.Errorf("resource %s not found for the principal %s", req.Resource, req.Principal)
	}

	// Get the evaluation for the specified action
	eval, evalExists := resourceMap[req.Action]
	if !evalExists {
		return Evaluation{}, fmt.Errorf("evaluation not found for the specified action %s on resource %s for principal %s", req.Action, req.Resource, req.Principal)
	}

	return eval, nil
}
