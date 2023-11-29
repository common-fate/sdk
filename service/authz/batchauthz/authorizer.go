package batchauthz

import (
	"context"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz"
)

type Authorizer interface {
	AddRequest(req authz.Request) Authorizer
	Execute(ctx context.Context, executor Executor) error
	IsPermitted(req authz.Request) (Evaluation, error)
}

type Executor interface {
	BatchAuthorize(context.Context, *connect.Request[authzv1alpha1.BatchAuthorizeRequest]) (*connect.Response[authzv1alpha1.BatchAuthorizeResponse], error)
}
