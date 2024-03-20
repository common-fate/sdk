package authz

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
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

func (c *Client) BatchAuthorize(ctx context.Context, input BatchAuthorizeInput) (BatchAuthorizeResponse, error) {
	r := make([]*authzv1alpha1.Request, 0)

	for _, item := range input.Requests {
		req, err := item.ToAPI()
		if err != nil {
			return BatchAuthorizeResponse{}, err
		}
		r = append(r, req)
	}

	res, err := c.raw.BatchAuthorize(ctx, connect.NewRequest(&authzv1alpha1.BatchAuthorizeRequest{
		Requests: r,
	}))
	if err != nil {
		return BatchAuthorizeResponse{}, err
	}

	ret := BatchAuthorizeResponse{
		Evaluations: res.Msg.Evaluations,
	}

	return ret, nil
}
