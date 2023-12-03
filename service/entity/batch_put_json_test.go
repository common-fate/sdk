package entity

import (
	"testing"

	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/stretchr/testify/assert"
)

func Test_transformJSONToEntity(t *testing.T) {
	tests := []struct {
		name         string
		give         EntityJSON
		want         *entityv1alpha1.Entity
		wantChildren []*entityv1alpha1.ChildRelation
		wantErr      bool
	}{
		{
			name: "ok",
			give: EntityJSON{
				UID: uid.UID{
					Type: "User",
					ID:   "test",
				},
				Attrs: map[string]any{
					"something": "string",
					"other":     1,
				},
			},
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "something",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Str{
								Str: "string",
							},
						},
					},
					{
						Key: "other",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Long{
								Long: 1,
							},
						},
					},
				},
			},
		},
		{
			name: "with set of entity refs",
			give: EntityJSON{
				UID: uid.UID{
					Type: "User",
					ID:   "test",
				},
				Attrs: map[string]any{
					"example": []any{
						map[string]any{
							"__entity": map[string]string{
								"type": "Foo",
								"id":   "foo",
							},
						},
					},
				},
			},
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "example",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Set{
								Set: &entityv1alpha1.Set{
									Values: []*entityv1alpha1.Value{
										{
											Value: &entityv1alpha1.Value_Entity{
												Entity: &entityv1alpha1.UID{
													Type: "Foo",
													Id:   "foo",
												},
											},
										},
									},
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
			got, children, err := transformJSONToEntity(tt.give)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformJSONToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantChildren, children)
		})
	}
}
