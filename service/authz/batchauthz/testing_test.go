package batchauthz

import (
	"context"
	"testing"

	"github.com/common-fate/sdk/service/authz"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/stretchr/testify/assert"
)

func TestMockBatchWorks(t *testing.T) {
	a := NewMock(t)

	req := authz.Request{
		Principal: uid.UID{
			Type: "User",
			ID:   "test",
		},
		Resource: uid.UID{
			Type: "Resource",
			ID:   "foo",
		},
		Action: uid.UID{
			Type: "Action",
			ID:   "bar",
		},
	}

	eval := Evaluation{
		ID:      "123",
		Allowed: true,
	}

	a.Mock(req, eval)

	a.Execute(context.Background())

	got, err := a.IsPermitted(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}
