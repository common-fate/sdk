package uid

import (
	"fmt"
	"strings"

	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/pkg/errors"
)

// UID is a unique identifier for an entity
// which includes both it's type and ID.
//
// e.g. 'GCP::Project::dev'
type UID struct {
	// Type of the entity.
	// e.g. in GCP::Project::dev
	// it is 'GCP::Project'
	Type string `json:"type"`

	// ID of the entity.
	// e.g. in GCP::Project::dev
	// it is 'dev'
	ID string `json:"id"`
}

// New creates a UID from a provided entity type and ID.
func New(entityType string, id string) UID {
	return UID{Type: entityType, ID: id}
}

func (u UID) ToAPI() *authzv1alpha1.UID {
	return &authzv1alpha1.UID{
		Type: u.Type,
		Id:   u.ID,
	}
}

// Valid returns true if the UID has both a type and an ID.
func (u UID) Valid() error {
	if u.Type == "" && u.ID == "" {
		return errors.New("UID was empty")
	}
	if u.Type == "" {
		return fmt.Errorf("UID with ID %s had no type", u.ID)
	}
	if u.ID == "" {
		return fmt.Errorf("UID with type %s had no ID", u.Type)
	}
	return nil
}

func (u UID) String() string {
	return fmt.Sprintf(`%s::"%s"`, u.Type, u.ID)
}

func FromAPI(uid *authzv1alpha1.UID) UID {
	if uid == nil {
		return UID{}
	}

	return UID{
		Type: uid.Type,
		ID:   uid.Id,
	}
}

func (u UID) Equals(other UID) bool {
	return u.Type == other.Type && u.ID == other.ID
}

// Parse extracts a UID from the input string.
// Unlike Cedar policies, we accept UIDs without mandatory double quotes
// around the ID component.
//
// For example: AWS::Account::123456789012 is a valid Common Fate UID, but
// is an invalid Cedar UID.
//
// We are more relaxed about parsing because nearly all of our UIDs do not
// contain any whitespace, so it saves users having to type additional
// quote characters when using our CLI.
//
// If the ID component of the UID contains any whitespace it must be
// enclosed in double quotes.
func Parse(input string) (UID, error) {
	parts := strings.Split(input, "::")

	// Check if there are at least two parts (type and ID)
	if len(parts) < 2 {
		return UID{}, errors.New("invalid UID format")
	}

	// Join the type parts with "::" separator
	// and assign it to the Type field
	uidType := strings.Join(parts[:len(parts)-1], "::")

	// Check if the type contains whitespace
	if len(strings.Fields(uidType)) > 1 {
		return UID{}, errors.New("whitespace not allowed in the UID type")
	}

	// The last part is the ID
	uidID := parts[len(parts)-1]

	// Check if the ID has spaces but is not enclosed in quotes
	if strings.Contains(uidID, " ") && !strings.HasPrefix(uidID, "\"") && !strings.HasSuffix(uidID, "\"") {
		return UID{}, errors.New("spaces in ID must be enclosed in quotes")
	}

	// If the ID is enclosed in quotes, remove them
	if strings.HasPrefix(uidID, "\"") && strings.HasSuffix(uidID, "\"") {
		uidID = uidID[1 : len(uidID)-1]
	}

	return UID{
		Type: uidType,
		ID:   uidID,
	}, nil
}

// ToSpecifier returns a Specifier which will match the UID.
func (u UID) ToSpecifier() *accessv1alpha1.Specifier {
	return &accessv1alpha1.Specifier{
		Specify: &accessv1alpha1.Specifier_Uid{
			Uid: u.ToAPI(),
		},
	}
}