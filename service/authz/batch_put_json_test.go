package authz

import (
	"testing"

	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func Test_transformJSONToEntity(t *testing.T) {
	tests := []struct {
		name    string
		give    EntityJSON
		want    *authzv1alpha1.Entity
		wantErr bool
	}{
		{
			name: "ok",
			give: EntityJSON{
				UID: UID{
					Type: "User",
					ID:   "test",
				},
				Attrs: map[string]any{
					"something": "string",
					"other":     1,
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "something",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Str{
								Str: "string",
							},
						},
					},
					{
						Key: "other",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Long{
								Long: 1,
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{},
			},
		},
		{
			name: "with set of entity refs",
			give: EntityJSON{
				UID: UID{
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
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "example",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Set{
								Set: &authzv1alpha1.Set{
									Values: []*authzv1alpha1.Value{
										{
											Value: &authzv1alpha1.Value_Entity{
												Entity: &authzv1alpha1.UID{
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
				Parents: []*authzv1alpha1.UID{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := transformJSONToEntity(tt.give)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformJSONToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
