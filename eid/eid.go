package eid

import (
	"encoding/json"
	"fmt"
	"strings"

	accessv1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/pkg/errors"
)

// EID is a unique identifier for an entity
// which includes both it's type and ID.
//
// e.g. 'GCP::Project::dev'
type EID struct {
	// Type of the entity.
	// e.g. in GCP::Project::dev
	// it is 'GCP::Project'
	Type string `json:"type"`

	// ID of the entity.
	// e.g. in GCP::Project::dev
	// it is 'dev'
	ID string `json:"id"`
}
type eidJson struct {
	Entity *EID `json:"__entity"`
}

func (e *EID) MarshalJSON() ([]byte, error) {
	return json.Marshal(eidJson{
		Entity: e,
	})
}

func (e *EID) UnmarshalJSON(data []byte) error {
	var aux eidJson
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	e.ID = aux.Entity.ID
	e.Type = aux.Entity.Type
	return nil
}

// New creates a EID from a provided entity type and ID.
func New(entityType string, id string) EID {
	return EID{Type: entityType, ID: id}
}

// Ptr creates a EID from a provided entity type and ID
// and returns it as a pointer.
func Ptr(entityType string, id string) *EID {
	return &EID{Type: entityType, ID: id}
}

func (u EID) ToAPI() *entityv1alpha1.EID {
	return &entityv1alpha1.EID{
		Type: u.Type,
		Id:   u.ID,
	}
}

// Valid returns true if the EID has both a type and an ID.
func (u EID) Valid() error {
	if u.Type == "" && u.ID == "" {
		return errors.New("EID was empty")
	}
	if u.Type == "" {
		return fmt.Errorf("EID with ID %s had no type", u.ID)
	}
	if u.ID == "" {
		return fmt.Errorf("EID with type %s had no ID", u.Type)
	}
	return nil
}

func (u EID) String() string {
	return fmt.Sprintf(`%s::"%s"`, u.Type, u.ID)
}

func FromAPI(eid *entityv1alpha1.EID) EID {
	if eid == nil {
		return EID{}
	}

	return EID{
		Type: eid.Type,
		ID:   eid.Id,
	}
}

func (u EID) Equals(other EID) bool {
	return u.Type == other.Type && u.ID == other.ID
}

// splitEID splits the string on "::" with specific conditions:
//   - Do not split if there is a single colon ':' before the double colon '::'.
//   - If there is a triple colon ':::' or more, only consider splitting on a double colon '::' that comes before it,
//     while still respecting the first rule.
func splitEID(s string) []string {
	var (
		colons       strings.Builder // Accumulates colon characters
		lastPosition int             // Tracks the start of the next component
		components   []string        // Holds the split components
	)

	var foundDouble bool // Flag to track if a valid double colon "::" has been found

	for i, char := range s {
		if char == ':' {
			colons.WriteRune(char) // Accumulate colon characters
		} else {
			colonSequence := colons.String() // Extract the accumulated colon sequence

			// If the sequence is a single colon and no double colon has been found,
			// return the entire string as it does not meet the split criteria.
			if colonSequence == ":" && !foundDouble {
				return []string{s}
			}

			// Check if the colon sequence is exactly "::"
			if colonSequence == "::" {
				foundDouble = true // Mark that a valid double colon has been found
				// Add the component found before this sequence (excluding single colons before it)
				components = append(components, s[lastPosition:i-len(colonSequence)])
				lastPosition = i // Update the start position for the next component
			}

			colons.Reset() // Reset the colon accumulator for the next sequence
		}
	}

	// Handle the last component if any colons are left
	colonSequence := colons.String()
	// Add the final component if the last colon sequence is less than or equal to "::"
	if len(colonSequence) <= 2 {
		components = append(components, s[lastPosition:])
	}

	return components
}

// Parse extracts a EID from the input string.
// Unlike Cedar policies, we accept EIDs without mandatory double quotes
// around the ID component.
//
// For example: AWS::Account::123456789012 is a valid Common Fate EID, but
// is an invalid Cedar EID.
//
// We are more relaxed about parsing because nearly all of our EIDs do not
// contain any whitespace, so it saves users having to type additional
// quote characters when using our CLI.
//
// If the ID component of the EID contains any whitespace it must be
// enclosed in double quotes.
func Parse(input string) (EID, error) {
	parts := splitEID(input)

	// Check if there are at least two parts (type and ID)
	if len(parts) < 2 {
		return EID{}, errors.New("invalid EID format")
	}

	// Join the type parts with "::" separator
	// and assign it to the Type field
	eidType := strings.Join(parts[:len(parts)-1], "::")

	// Check if the type contains whitespace
	if len(strings.Fields(eidType)) > 1 {
		return EID{}, errors.New("whitespace not allowed in the EID type")
	}

	// The last part is the ID
	eidID := parts[len(parts)-1]

	// Check if the ID has spaces but is not enclosed in quotes
	if strings.Contains(eidID, " ") && !strings.HasPrefix(eidID, "\"") && !strings.HasSuffix(eidID, "\"") {
		return EID{}, errors.New("spaces in ID must be enclosed in quotes")
	}

	// If the ID is enclosed in quotes, remove them
	if strings.HasPrefix(eidID, "\"") && strings.HasSuffix(eidID, "\"") {
		eidID = eidID[1 : len(eidID)-1]
	}

	return EID{
		Type: eidType,
		ID:   eidID,
	}, nil
}

// ToSpecifier returns a Specifier which will match the EID.
func (u EID) ToSpecifier() *accessv1alpha1.Specifier {
	return &accessv1alpha1.Specifier{
		Specify: &accessv1alpha1.Specifier_Eid{
			Eid: u.ToAPI(),
		},
	}
}
