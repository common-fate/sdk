package entity

import (
	"reflect"
	"testing"
	"time"

	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/stretchr/testify/assert"
)

type testUser struct {
	ID   string `authz:"id"`
	Name string `authz:"name"`
}

func (t testUser) UID() uid.UID { return uid.New("User", t.ID) }

type testAccount struct {
	ID      string `authz:"id"`
	Name    string `authz:"name"`
	OrgUnit string
}

func (t testAccount) Parents() []uid.UID { return []uid.UID{uid.New("OrgUnit", t.OrgUnit)} }
func (t testAccount) UID() uid.UID       { return uid.New("Account", t.ID) }

type testVault struct {
	ID        string `authz:"id"`
	LongVal   int    `authz:"long_val"`
	BoolVal   bool   `authz:"bool_val"`
	StringVal string `authz:"string_val"`
}

func (t testVault) UID() uid.UID { return uid.New("Vault", t.ID) }

type testAnyParents struct {
	ID        string    `authz:"id"`
	Resources []uid.UID // no tag here as we just want the parent relations, we don't want this as an attr
}

func (t testAnyParents) Parents() []uid.UID { return t.Resources }
func (t testAnyParents) UID() uid.UID       { return uid.New("AnyParent", t.ID) }

type testAccessRequest struct {
	ID       string  `authz:"id"`
	Resource uid.UID `authz:"resource"`
}

func (t testAccessRequest) Parents() []uid.UID { return []uid.UID{t.Resource} }
func (t testAccessRequest) UID() uid.UID       { return uid.New("AccessRequest", t.ID) }

type testEntitlement struct {
	ID     string  `authz:"id"`
	Target uid.UID `authz:"target"`
}

func (t testEntitlement) UID() uid.UID     { return uid.New("Entitlement", t.ID) }
func (testEntitlement) EntityType() string { return "Entitlement" }

type testEntityWithSet struct {
	ID     string   `authz:"id"`
	Things []string `authz:"things"`
	Other  []int    `authz:"other"`
}

func (t testEntityWithSet) UID() uid.UID     { return uid.New("WithSet", t.ID) }
func (testEntityWithSet) EntityType() string { return "WithSet" }

type testEntityWithSetOfReferences struct {
	ID   string    `authz:"id"`
	Refs []uid.UID `authz:"refs"`
}

func (t testEntityWithSetOfReferences) UID() uid.UID     { return uid.New("WithSetOfReferences", t.ID) }
func (testEntityWithSetOfReferences) EntityType() string { return "WithSetOfReferences" }

type exampleRecord struct {
	Foo string `authz:"foo"`
	Bar int    `authz:"bar"`
}

type testEntityWithRecord struct {
	ID      string        `authz:"id"`
	Example exampleRecord `authz:"example_record"`
}

func (t testEntityWithRecord) UID() uid.UID     { return uid.New("WithRecord", t.ID) }
func (testEntityWithRecord) EntityType() string { return "WithRecord" }

type entityRefs struct {
	Foo uid.UID `authz:"foo"`
	Bar uid.UID `authz:"bar"`
}

type testEntityWithRecordOfRefs struct {
	ID   string     `authz:"id"`
	Refs entityRefs `authz:"refs"`
}

func (t testEntityWithRecordOfRefs) UID() uid.UID { return uid.New("WithRecordOfRefs", t.ID) }

type recordOfSets struct {
	Foo    []string      `authz:"foo"`
	Bar    []int         `authz:"bar"`
	Nested exampleRecord `authz:"nested"`
}

type testEntityWithRecordOfSets struct {
	ID     string       `authz:"id"`
	Record recordOfSets `authz:"record"`
}

func (t testEntityWithRecordOfSets) UID() uid.UID { return uid.New("WithRecordOfSets", t.ID) }

type testEntityWithParentsMethod struct {
	ID     string  `authz:"id"`
	Target uid.UID `authz:"target"`
	Other  uid.UID `authz:"other"`
}

func (t testEntityWithParentsMethod) UID() uid.UID       { return uid.New("WithParentsMethod", t.ID) }
func (e testEntityWithParentsMethod) Parents() []uid.UID { return []uid.UID{e.Other, e.Target} }

type testEntityWithOptionalField struct {
	ID       string         `authz:"id"`
	Foo      *string        `authz:"foo"`
	Bar      *uid.UID       `authz:"bar"`
	Long     *int           `authz:"long"`
	Time     *time.Time     `authz:"time"`
	Duration *time.Duration `authz:"duration"`
}

func (t testEntityWithOptionalField) UID() uid.UID { return uid.New("WithOptionalField", t.ID) }

type testEntityWithTimeField struct {
	ID        string        `authz:"id"`
	CreatedAt time.Time     `authz:"created_at"`
	Other     time.Duration `authz:"other"`
}

func (t testEntityWithTimeField) UID() uid.UID { return uid.New("WithTimeField", t.ID) }

func Test_transformToEntity(t *testing.T) {
	type args struct {
		e Entity
	}
	tests := []struct {
		name         string
		args         args
		want         *entityv1alpha1.Entity
		wantChildren []*entityv1alpha1.ChildRelation
		wantErr      bool
	}{
		{
			name: "ok",
			args: args{
				e: testUser{
					ID:   "test",
					Name: "testing",
				},
			},
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "User",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "name",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Str{
								Str: "testing",
							},
						},
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
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "Account",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "name",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Str{
								Str: "testing",
							},
						},
					},
				},
			},
			wantChildren: []*entityv1alpha1.ChildRelation{
				{
					Parent: &entityv1alpha1.UID{
						Type: "OrgUnit",
						Id:   "prod",
					},
					Child: &entityv1alpha1.UID{
						Type: "Account",
						Id:   "test",
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
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "Vault",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "long_val",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Long{
								Long: 1,
							},
						},
					},
					{
						Key: "bool_val",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Bool{
								Bool: true,
							},
						},
					},
					{
						Key: "string_val",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Str{
								Str: "hi",
							},
						},
					},
				},
			},
		},
		{
			name: "generic parents",
			args: args{
				e: testAnyParents{
					ID: "test",
					Resources: []uid.UID{
						{
							Type: "Something",
							ID:   "test",
						},
						{
							Type: "Other",
							ID:   "else",
						},
					},
				},
			},
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "AnyParent",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{},
			},
			wantChildren: []*entityv1alpha1.ChildRelation{
				{
					Parent: &entityv1alpha1.UID{
						Type: "Something",
						Id:   "test",
					},
					Child: &entityv1alpha1.UID{
						Type: "AnyParent",
						Id:   "test",
					},
				},
				{
					Parent: &entityv1alpha1.UID{
						Type: "Other",
						Id:   "else",
					},
					Child: &entityv1alpha1.UID{
						Type: "AnyParent",
						Id:   "test",
					},
				},
			},
		},
		{
			name: "access request",
			args: args{
				e: testAccessRequest{
					ID:       "test",
					Resource: uid.New("Something", "test"),
				},
			},
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "AccessRequest",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "resource",

						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Entity{
								Entity: uid.New("Something", "test").ToAPI(),
							},
						},
					},
				},
			},
			wantChildren: []*entityv1alpha1.ChildRelation{
				{
					Parent: &entityv1alpha1.UID{
						Type: "Something",
						Id:   "test",
					},
					Child: &entityv1alpha1.UID{
						Type: "AccessRequest",
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
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "Entitlement",
					Id:   "123",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "target",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Entity{
								Entity: &entityv1alpha1.UID{
									Type: "Vault",
									Id:   "test",
								},
							},
						},
					},
				},
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
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "WithParentsMethod",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "target",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Entity{
								Entity: &entityv1alpha1.UID{
									Type: "Foo",
									Id:   "1",
								},
							},
						},
					},
					{
						Key: "other",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Entity{
								Entity: &entityv1alpha1.UID{
									Type: "Bar",
									Id:   "2",
								},
							},
						},
					},
				},
			},
			wantChildren: []*entityv1alpha1.ChildRelation{
				{
					Parent: &entityv1alpha1.UID{
						Type: "Bar",
						Id:   "2",
					},
					Child: &entityv1alpha1.UID{
						Type: "WithParentsMethod",
						Id:   "test",
					},
				},
				{
					Parent: &entityv1alpha1.UID{
						Type: "Foo",
						Id:   "1",
					},
					Child: &entityv1alpha1.UID{
						Type: "WithParentsMethod",
						Id:   "test",
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
			want: &entityv1alpha1.Entity{
				Uid: &entityv1alpha1.UID{
					Type: "WithTimeField",
					Id:   "test",
				},
				Attributes: []*entityv1alpha1.Attribute{
					{
						Key: "created_at",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Long{
								Long: exampleTime.UnixMilli(),
							},
						},
					},
					{
						Key: "other",
						Value: &entityv1alpha1.Value{
							Value: &entityv1alpha1.Value_Long{
								Long: 1000,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, children, err := EntityToAPI(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformToEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantChildren, children)
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
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				input: `authz:"id"`,
			},
			want: "id",
		},
		{
			name: "parent",
			args: args{
				input: `authz:"group,parent=Group"`,
			},
			want: "group",
		},
		{
			name: "with other tags",
			args: args{
				input: `authz:"group,parent=Group,something=something" json:"something"`,
			},
			want: "group",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseTag(tt.args.input)
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
			res, _, err := EntityToAPI(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			err = Unmarshal(res, tt.out)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.in, tt.out)
		})
	}
}
