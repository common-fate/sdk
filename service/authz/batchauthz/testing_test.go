package batchauthz

import (
	"context"
	"testing"

	"github.com/common-fate/sdk/service/authz"
	"github.com/common-fate/sdk/uid"
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

	a.Authorize(context.Background())

	got, err := a.IsPermitted(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}

func TestMockAnnotations(t *testing.T) {
	a := NewMock(t)

	principal := uid.New("User", "test")
	resource := uid.New("Resource", "foo")
	action := uid.New("Action", "bar")

	annos := NewAnnotations()
	annos.Set("foo", Annotation{Value: "bar"})
	annos.Set("foo", Annotation{Value: "other"})

	eval := Evaluation{
		Allowed:     true,
		Annotations: annos,
	}

	a.Allow(principal, resource, action).Annotations("foo", "bar", "other")

	a.Authorize(context.Background())

	got, err := a.IsPermitted(authz.Request{Principal: principal, Resource: resource, Action: action})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}
