package batchauthz

import (
	"context"
	"testing"

	"github.com/common-fate/sdk/service/authz"
	"github.com/stretchr/testify/assert"
)

func TestMockBatchWorks(t *testing.T) {
	a := NewMock(t)

	req := Request{
		Principal: authz.UID{
			Type: "User",
			ID:   "test",
		},
		Resource: authz.UID{
			Type: "Resource",
			ID:   "foo",
		},
		Action: authz.UID{
			Type: "Action",
			ID:   "bar",
		},
	}

	eval := Evaluation{
		ID:      "123",
		Allowed: true,
	}

	a.Mock(req, eval)

	a.Execute(context.Background(), nil)

	got, err := a.IsPermitted(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}
