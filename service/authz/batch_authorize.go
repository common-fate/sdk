package authz

import (
	"fmt"

	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type DecisionNotFoundError struct {
	Key string
}

func (e DecisionNotFoundError) Error() string {
	return fmt.Sprintf("decision not found for client_id: %s", e.Key)
}

type BatchAuthorizeInput struct {
	Requests map[string]Request
}

type BatchAuthorizeResponse struct {
	// Evaluations made by the authz service.
	Evaluations []*authzv1alpha1.Evaluation
}
