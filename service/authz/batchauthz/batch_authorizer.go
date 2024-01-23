package batchauthz

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz"
	"github.com/common-fate/sdk/service/entity"
)

// Batch is a batch authorization request.
// You should construct a Batch when you need to test multiple authorizations
// at once. Batches should be scoped to a particular API call and should not
// be shared between API calls.
type Batch struct {
	request *authzv1alpha1.BatchAuthorizeRequest
	// evaluations is a map of principal -> resource -> action -> Evaluation
	evaluations map[eid.EID]map[eid.EID]map[eid.EID]*authzv1alpha1.Evaluation
	executor    Executor
	executed    bool
}

var _ Authorizer = &Batch{}

func New(executor Executor) *Batch {
	return &Batch{
		request:     &authzv1alpha1.BatchAuthorizeRequest{},
		executor:    executor,
		evaluations: map[eid.EID]map[eid.EID]map[eid.EID]*authzv1alpha1.Evaluation{},
	}
}

func (a *Batch) AddRequest(req authz.Request) error {
	apiReq := &authzv1alpha1.Request{
		Principal: req.Principal.ToAPI(),
		Action:    req.Action.ToAPI(),
		Resource:  req.Resource.ToAPI(),
	}

	for _, e := range req.OverlayEntities {
		apiEntity, childRels, err := entity.Marshal(e)
		if err != nil {
			return err
		}

		apiReq.OverlayEntities = append(apiReq.OverlayEntities, apiEntity)
		apiReq.OverlayChildren = append(apiReq.OverlayChildren, childRels...)
	}

	for _, c := range req.OverlayChildren {
		apiReq.OverlayChildren = append(apiReq.OverlayChildren, c.ToAPI())
	}

	a.request.Requests = append(a.request.Requests, apiReq)

	return nil
}

func (a *Batch) Authorize(ctx context.Context) error {
	if a.executed {
		return errors.New("Authorize() has already been called: create a new batchauthz.Batch rather than reusing the same one multiple times")
	}

	res, err := a.executor.BatchAuthorize(ctx, connect.NewRequest(a.request))
	if err != nil {
		return err
	}

	for _, eval := range res.Msg.Evaluations {
		principal := eid.FromAPI(eval.Request.Principal)
		action := eid.FromAPI(eval.Request.Action)
		resource := eid.FromAPI(eval.Request.Resource)

		// Check if the principal exists in the evaluations map
		principalMap, ok := a.evaluations[principal]
		if !ok {
			principalMap = make(map[eid.EID]map[eid.EID]*authzv1alpha1.Evaluation)
			a.evaluations[principal] = principalMap
		}

		// Check if the resource map exists for the principal
		resourceMap, ok := principalMap[resource]
		if !ok {
			resourceMap = make(map[eid.EID]*authzv1alpha1.Evaluation)
			principalMap[resource] = resourceMap
		}

		existingEval, ok := resourceMap[action]
		if ok {
			// there were multiple evaluations for a particular principal/action/resource combination.
			// this is a logic error and so we return an error
			return fmt.Errorf("multiple authorization evaluations found for principal %s, action %s and resource %s (%s and %s): ensure that AddRequest is only called once for each combination of principal/action/resource", principal, action, resource, existingEval.Id, eval.Id)
		}

		// Populate the evaluations map
		resourceMap[action] = eval
	}

	a.executed = true

	return nil
}

type Evaluation struct {
	// The evaluation ID from authz
	ID string
	// Whether access is permitted
	Allowed bool
	// Policies which contributed to the authorization decision
	Policies []string
	// Annotations on the contributing policies
	Annotations Annotations
	// Errors from evaluating the policies (if any)
	Errors []string
}

type Annotation struct {
	Value    string
	PolicyID string
}

type Annotations struct {
	annotations map[string][]Annotation
}

func (a *Annotations) All() map[string][]Annotation {
	return a.annotations
}

func NewAnnotations() Annotations {
	return Annotations{annotations: map[string][]Annotation{}}
}

func (a *Annotations) Set(key string, annotation Annotation) {
	if a.annotations == nil {
		a.annotations = map[string][]Annotation{}
	}

	// Get the current list of annotations for the key
	annotationList, exists := a.annotations[key]

	// If the key doesn't exist, create a new entry
	if !exists {
		annotationList = []Annotation{}
	}

	// Append the new annotation to the list
	annotationList = append(annotationList, annotation)

	// Update the annotations map
	a.annotations[key] = annotationList
}

func (a Annotations) GetValues(key string) []string {
	vals, ok := a.annotations[key]
	if !ok {
		return nil
	}
	res := make([]string, len(vals))

	for i, val := range vals {
		res[i] = val.Value
	}

	return res
}

func (a *Batch) IsPermitted(req authz.Request) (Evaluation, error) {
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

	if eval.Decision == authzv1alpha1.Decision_DECISION_UNSPECIFIED {
		return Evaluation{}, errors.New("evaluation decision was unspecified")
	}

	// Map the authzv1alpha1.Evaluation to the Evaluation struct
	result := Evaluation{
		ID:          eval.Id,
		Allowed:     eval.Decision == authzv1alpha1.Decision_DECISION_ALLOW,
		Policies:    eval.Diagnostics.Reason,
		Annotations: NewAnnotations(),
		Errors:      eval.Diagnostics.Errors,
	}

	// Copy annotations
	for _, anno := range eval.Diagnostics.Annotations {
		result.Annotations.Set(anno.Key, Annotation{Value: anno.Value, PolicyID: anno.PolicyId})
	}

	return result, nil
}
