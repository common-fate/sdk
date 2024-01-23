package policyset

import (
	"context"

	"connectrpc.com/connect"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
)

type DeleteInput struct {
	ID string
}

type DeleteOutput struct {
	ID string
}

func (c *Client) Delete(ctx context.Context, input DeleteInput) (DeleteOutput, error) {
	apires, err := c.raw.DeletePolicySet(ctx, connect.NewRequest(&authzv1alpha1.DeletePolicySetRequest{
		Id: input.ID,
	}))
	if err != nil {
		return DeleteOutput{}, err
	}

	res := DeleteOutput{
		ID: apires.Msg.Id,
	}

	return res, nil
}
