package batchauthz

import (
	"context"
	"testing"

	"github.com/common-fate/sdk/eid"
	"github.com/common-fate/sdk/service/authz"
	"github.com/stretchr/testify/assert"
)

func TestMockBatchWorks(t *testing.T) {
	a := NewMock(t)

	req := authz.Request{
		Principal: eid.EID{
			Type: "User",
			ID:   "test",
		},
		Resource: eid.EID{
			Type: "Resource",
			ID:   "foo",
		},
		Action: eid.EID{
			Type: "Action",
			ID:   "bar",
		},
	}

	eval := Evaluation{
		ID:      "123",
		Allowed: true,
	}

	a.Mock(req, eval)

	err := a.Authorize(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	got, err := a.IsPermitted(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}

func TestMockAnnotations(t *testing.T) {
	a := NewMock(t)

	principal := eid.New("User", "test")
	resource := eid.New("Resource", "foo")
	action := eid.New("Action", "bar")

	annos := NewAnnotations()
	annos.Set("foo", Annotation{Value: "bar"})
	annos.Set("foo", Annotation{Value: "other"})

	eval := Evaluation{
		Allowed:     true,
		Annotations: annos,
	}

	a.Allow(principal, resource, action).Annotations("foo", "bar", "other")

	err := a.Authorize(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	got, err := a.IsPermitted(authz.Request{Principal: principal, Resource: resource, Action: action})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, eval, got)
}
