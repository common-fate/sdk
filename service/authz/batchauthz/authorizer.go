package batchauthz

import (
	"context"

	"connectrpc.com/connect"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz"
)

type Authorizer interface {
	// AddRequest adds a new authorization request to be evaluated.
	AddRequest(req authz.Request) error
	// Authorization calls the authz service to evaluate all of the authorization requests added.
	Authorize(ctx context.Context) error
	// IsPermitted returns the evaluation associated with the given request.
	// Ensure that you call AddRequest() and Authorize() before calling IsPermitted().
	IsPermitted(req authz.Request) (Evaluation, error)
}

type Executor interface {
	BatchAuthorize(context.Context, *connect.Request[authzv1alpha1.BatchAuthorizeRequest]) (*connect.Response[authzv1alpha1.BatchAuthorizeResponse], error)
}
