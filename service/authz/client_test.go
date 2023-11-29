package authz

import (
	"reflect"
	"testing"
	"time"

	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/stretchr/testify/assert"
)

type testUser struct {
	ID     string   `authz:"id"`
	Name   string   `authz:"name"`
	Groups []string `authz:"groups,parent=Group"`
}

func (testUser) EntityType() string { return "User" }

type testAccount struct {
	ID      string `authz:"id"`
	Name    string `authz:"name"`
	OrgUnit string `authz:"org_unit,parent=OrgUnit"`
}

func (testAccount) EntityType() string { return "Account" }

type testVault struct {
	ID        string `authz:"id"`
	LongVal   int    `authz:"long_val"`
	BoolVal   bool   `authz:"bool_val"`
	StringVal string `authz:"string_val"`
}

func (testVault) EntityType() string { return "Vault" }

type testAnyParents struct {
	ID        string               `authz:"id"`
	Resources []*authzv1alpha1.UID `authz:"resources,parent"`
}

func (testAnyParents) EntityType() string { return "AnyParent" }

type testAccessRequest struct {
	ID       string             `authz:"id"`
	Resource *authzv1alpha1.UID `authz:"resource,parent"`
}

func (testAccessRequest) EntityType() string { return "AccessRequest" }

type testEntitlement struct {
	ID     string  `authz:"id"`
	Target uid.UID `authz:"target"`
}

func (testEntitlement) EntityType() string { return "Entitlement" }

type testEntityWithSet struct {
	ID     string   `authz:"id"`
	Things []string `authz:"things"`
	Other  []int    `authz:"other"`
}

func (testEntityWithSet) EntityType() string { return "WithSet" }

type testEntityWithSetOfReferences struct {
	ID   string    `authz:"id"`
	Refs []uid.UID `authz:"refs"`
}

func (testEntityWithSetOfReferences) EntityType() string { return "WithSetOfReferences" }

type exampleRecord struct {
	Foo string `authz:"foo"`
	Bar int    `authz:"bar"`
}

type testEntityWithRecord struct {
	ID      string        `authz:"id"`
	Example exampleRecord `authz:"example_record"`
}

func (testEntityWithRecord) EntityType() string { return "WithRecord" }

type entityRefs struct {
	Foo uid.UID `authz:"foo"`
	Bar uid.UID `authz:"bar"`
}

type testEntityWithRecordOfRefs struct {
	ID   string     `authz:"id"`
	Refs entityRefs `authz:"refs"`
}

func (testEntityWithRecordOfRefs) EntityType() string { return "WithRecordOfRefs" }

type recordOfSets struct {
	Foo    []string      `authz:"foo"`
	Bar    []int         `authz:"bar"`
	Nested exampleRecord `authz:"nested"`
}

type testEntityWithRecordOfSets struct {
	ID     string       `authz:"id"`
	Record recordOfSets `authz:"record"`
}

func (testEntityWithRecordOfSets) EntityType() string { return "WithRecordOfSets" }

type testEntityWithParentsMethod struct {
	ID     string  `authz:"id"`
	Target uid.UID `authz:"target"`
	Other  uid.UID `authz:"other"`
}

func (testEntityWithParentsMethod) EntityType() string   { return "WithParentsMethod" }
func (e testEntityWithParentsMethod) Parents() []uid.UID { return []uid.UID{e.Other, e.Target} }

type testEntityWithOptionalField struct {
	ID       string         `authz:"id"`
	Foo      *string        `authz:"foo"`
	Bar      *uid.UID       `authz:"bar"`
	Long     *int           `authz:"long"`
	Time     *time.Time     `authz:"time"`
	Duration *time.Duration `authz:"duration"`
}

func (testEntityWithOptionalField) EntityType() string { return "WithOptionalField" }

type testEntityWithTimeField struct {
	ID        string        `authz:"id"`
	CreatedAt time.Time     `authz:"created_at"`
	Other     time.Duration `authz:"other"`
}

func (testEntityWithTimeField) EntityType() string { return "WithTimeField" }

func Test_transformToEntity(t *testing.T) {
	type args struct {
		e Entity
	}
	tests := []struct {
		name    string
		args    args
		want    *authzv1alpha1.Entity
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				e: testUser{
					ID:     "test",
					Name:   "testing",
					Groups: []string{"devs"},
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "name",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Str{
								Str: "testing",
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{
					{
						Type: "Group",
						Id:   "devs",
					},
				},
			},
		},
		{
			name: "many to one",
			args: args{
				e: testAccount{
					ID:      "test",
					Name:    "testing",
					OrgUnit: "prod",
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "Account",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "name",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Str{
								Str: "testing",
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{
					{
						Type: "OrgUnit",
						Id:   "prod",
					},
				},
			},
		},
		{
			name: "attribute parsing",
			args: args{
				e: testVault{
					ID:        "test",
					LongVal:   1,
					BoolVal:   true,
					StringVal: "hi",
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "Vault",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "long_val",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Long{
								Long: 1,
							},
						},
					},
					{
						Key: "bool_val",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Bool{
								Bool: true,
							},
						},
					},
					{
						Key: "string_val",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Str{
								Str: "hi",
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{},
			},
		},
		{
			name: "generic parents",
			args: args{
				e: testAnyParents{
					ID: "test",
					Resources: []*authzv1alpha1.UID{
						{
							Type: "Something",
							Id:   "test",
						},
						{
							Type: "Other",
							Id:   "else",
						},
					},
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "AnyParent",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{},
				Parents: []*authzv1alpha1.UID{
					{
						Type: "Something",
						Id:   "test",
					},
					{
						Type: "Other",
						Id:   "else",
					},
				},
			},
		},
		{
			name: "access request",
			args: args{
				e: testAccessRequest{
					ID: "test",
					Resource: &authzv1alpha1.UID{
						Type: "Something",
						Id:   "test",
					},
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "AccessRequest",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{},
				Parents: []*authzv1alpha1.UID{
					{
						Type: "Something",
						Id:   "test",
					},
				},
			},
		},
		{
			name: "entity reference as field",
			args: args{
				e: testEntitlement{
					ID: "123",
					Target: uid.UID{
						Type: "Vault",
						ID:   "test",
					},
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "Entitlement",
					Id:   "123",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "target",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Entity{
								Entity: &authzv1alpha1.UID{
									Type: "Vault",
									Id:   "test",
								},
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{},
			},
		},
		{
			name: "parents from Parents() method",
			args: args{
				e: testEntityWithParentsMethod{
					ID: "test",
					Target: uid.UID{
						Type: "Foo",
						ID:   "1",
					},
					Other: uid.UID{
						Type: "Bar",
						ID:   "2",
					},
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "WithParentsMethod",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "target",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Entity{
								Entity: &authzv1alpha1.UID{
									Type: "Foo",
									Id:   "1",
								},
							},
						},
					},
					{
						Key: "other",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Entity{
								Entity: &authzv1alpha1.UID{
									Type: "Bar",
									Id:   "2",
								},
							},
						},
					},
				},
				Parents: []*authzv1alpha1.UID{
					{
						Type: "Bar",
						Id:   "2",
					},
					{
						Type: "Foo",
						Id:   "1",
					},
				},
			},
		},
		{
			name: "time and duration",
			args: args{
				e: testEntityWithTimeField{
					ID:        "test",
					CreatedAt: exampleTime,
					Other:     time.Second,
				},
			},
			want: &authzv1alpha1.Entity{
				Uid: &authzv1alpha1.UID{
					Type: "WithTimeField",
					Id:   "test",
				},
				Attributes: []*authzv1alpha1.Attribute{
					{
						Key: "created_at",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Long{
								Long: exampleTime.UnixMilli(),
							},
						},
					},
					{
						Key: "other",
						Value: &authzv1alpha1.Value{
							Value: &authzv1alpha1.Value_Long{
								Long: 1000,
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
			got, err := transformToEntity(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

var exampleString = "EXAMPLE"
var exampleInt = 10
var exampleTime = time.Now().Round(time.Millisecond)
var exampleDuration = time.Second

func Test_parseTag(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    Tag
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				input: `authz:"id"`,
			},
			want: Tag{
				Name: "id",
			},
		},
		{
			name: "parent",
			args: args{
				input: `authz:"group,parent=Group"`,
			},
			want: Tag{
				Name:       "group",
				ParentType: "Group",
				HasParent:  true,
			},
		},
		{
			name: "with other tags",
			args: args{
				input: `authz:"group,parent=Group" json:"something"`,
			},
			want: Tag{
				Name:       "group",
				ParentType: "Group",
				HasParent:  true,
			},
		},
		{
			name: "with generic parent",
			args: args{
				input: `authz:"group,parent"`,
			},
			want: Tag{
				Name:      "group",
				HasParent: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTag(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

// we should be able to roundtrip from/to an entity into the protobuf definition and back.
func TestUnmarshalEntity_roundtrip(t *testing.T) {
	tests := []struct {
		name string
		in   Entity
		out  Entity
	}{
		{
			name: "basic entity",
			in: &testVault{
				ID:        "123",
				LongVal:   10,
				BoolVal:   true,
				StringVal: "hi",
			},
			out: &testVault{},
		},
		{
			name: "entity reference as a field",
			in: &testEntitlement{
				ID: "123",
				Target: uid.UID{
					Type: "Vault",
					ID:   "test",
				},
			},
			out: &testEntitlement{},
		},
		{
			name: "entity with set",
			in: &testEntityWithSet{
				ID:     "123",
				Things: []string{"hello", "world"},
				Other:  []int{1, 2},
			},
			out: &testEntityWithSet{},
		},
		{
			name: "entity with set of entity references",
			in: &testEntityWithSetOfReferences{
				ID: "123",
				Refs: []uid.UID{
					{
						Type: "Something",
						ID:   "1",
					},
					{
						Type: "Other",
						ID:   "2",
					},
				},
			},
			out: &testEntityWithSetOfReferences{},
		},
		{
			name: "record attr",
			in: &testEntityWithRecord{
				ID: "123",
				Example: exampleRecord{
					Foo: "hello",
					Bar: 1,
				},
			},
			out: &testEntityWithRecord{},
		},
		{
			name: "record of entity refs",
			in: &testEntityWithRecordOfRefs{
				ID: "123",
				Refs: entityRefs{
					Foo: uid.UID{
						Type: "Something",
						ID:   "1",
					},
					Bar: uid.UID{
						Type: "Other",
						ID:   "2",
					},
				},
			},
			out: &testEntityWithRecordOfRefs{},
		},
		{
			name: "record of sets and records",
			in: &testEntityWithRecordOfSets{
				ID: "123",
				Record: recordOfSets{
					Foo: []string{"hello", "world"},
					Bar: []int{1, 2, 3},
					Nested: exampleRecord{
						Foo: "hi",
						Bar: 10,
					},
				},
			},
			out: &testEntityWithRecordOfSets{},
		},
		{
			name: "optional field provided",
			in: &testEntityWithOptionalField{
				ID:  "123",
				Foo: &exampleString,
				Bar: &uid.UID{
					Type: "Something",
					ID:   "1",
				},
				Time:     &exampleTime,
				Long:     &exampleInt,
				Duration: &exampleDuration,
			},
			out: &testEntityWithOptionalField{},
		},
		{
			name: "optional field empty",
			in: &testEntityWithOptionalField{
				ID:  "123",
				Foo: nil,
			},
			out: &testEntityWithOptionalField{},
		},
		{
			name: "with time field",
			in: &testEntityWithTimeField{
				ID:        "123",
				CreatedAt: time.Now().Round(time.Millisecond),
				Other:     time.Hour,
			},
			out: &testEntityWithTimeField{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := transformToEntity(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			err = UnmarshalEntity(res, tt.out)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.in, tt.out)
		})
	}
}
