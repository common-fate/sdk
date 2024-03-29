package batchauthz

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/common-fate/sdk/eid"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/service/authz"
	"github.com/stretchr/testify/assert"
)

// mockExecutor is a mock implementation of the Executor interface.
type mockExecutor struct {
	Response *authzv1alpha1.BatchAuthorizeResponse
}

// BatchAuthorize is a mocked implementation of the BatchAuthorize method.
func (m *mockExecutor) BatchAuthorize(ctx context.Context, req *connect.Request[authzv1alpha1.BatchAuthorizeRequest]) (*connect.Response[authzv1alpha1.BatchAuthorizeResponse], error) {
	return connect.NewResponse(m.Response), nil
}

func TestBatch_Authorize(t *testing.T) {
	tests := []struct {
		name            string
		executor        mockExecutor
		wantExecuted    bool
		wantEvaluations map[eid.EID]map[eid.EID]map[eid.EID]*authzv1alpha1.Evaluation
		wantErr         bool
	}{
		{
			name: "ok",
			executor: mockExecutor{
				Response: &authzv1alpha1.BatchAuthorizeResponse{
					Evaluations: []*authzv1alpha1.Evaluation{
						{
							Id: "eval_1",
							Request: &authzv1alpha1.Request{
								Principal: &entityv1alpha1.EID{
									Type: "User",
									Id:   "test",
								},
								Action: &entityv1alpha1.EID{
									Type: "Action",
									Id:   "foo",
								},
								Resource: &entityv1alpha1.EID{
									Type: "Resource",
									Id:   "bar",
								},
							},
						},
					},
				},
			},
			wantExecuted: true,
			wantEvaluations: map[eid.EID]map[eid.EID]map[eid.EID]*authzv1alpha1.Evaluation{
				{Type: "User", ID: "test"}: {
					{Type: "Resource", ID: "bar"}: {
						{Type: "Action", ID: "foo"}: {
							Id: "eval_1",
							Request: &authzv1alpha1.Request{
								Principal: &entityv1alpha1.EID{
									Type: "User",
									Id:   "test",
								},
								Action: &entityv1alpha1.EID{
									Type: "Action",
									Id:   "foo",
								},
								Resource: &entityv1alpha1.EID{
									Type: "Resource",
									Id:   "bar",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(&tt.executor)
			if err := a.Authorize(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Batch.Authorize() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantEvaluations, a.evaluations)
			assert.Equal(t, tt.wantExecuted, a.executed)
		})
	}
}

func TestBatch_IsPermitted(t *testing.T) {
	testAnno := NewAnnotations()
	testAnno.Set("advice", Annotation{
		Value:    "something",
		PolicyID: "policy1",
	})

	tests := []struct {
		name     string
		executor mockExecutor
		give     authz.Request
		wantEval Evaluation
	}{
		{
			name: "ok",
			executor: mockExecutor{
				Response: &authzv1alpha1.BatchAuthorizeResponse{
					Evaluations: []*authzv1alpha1.Evaluation{
						{
							Id:       "eval_1",
							Decision: authzv1alpha1.Decision_DECISION_ALLOW,
							Diagnostics: &authzv1alpha1.Diagnostics{
								Reason: []string{"policy1"},
								Errors: []string{"error1"},
								Annotations: []*authzv1alpha1.Annotation{
									{
										Key:      "advice",
										PolicyId: "policy1",
										Value:    "something",
									},
								},
							},
							Request: &authzv1alpha1.Request{
								Principal: &entityv1alpha1.EID{
									Type: "User",
									Id:   "test",
								},
								Action: &entityv1alpha1.EID{
									Type: "Action",
									Id:   "foo",
								},
								Resource: &entityv1alpha1.EID{
									Type: "Resource",
									Id:   "bar",
								},
							},
						},
					},
				},
			},
			wantEval: Evaluation{
				ID:          "eval_1",
				Allowed:     true,
				Policies:    []string{"policy1"},
				Annotations: testAnno,
				Errors:      []string{"error1"},
			},
			give: authz.Request{
				Principal: eid.EID{
					Type: "User",
					ID:   "test",
				},
				Action: eid.EID{
					Type: "Action",
					ID:   "foo",
				},
				Resource: eid.EID{
					Type: "Resource",
					ID:   "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(&tt.executor)
			err := a.Authorize(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			got, err := a.IsPermitted(tt.give)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.wantEval, got)
		})
	}
}
