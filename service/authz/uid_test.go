package authz

import (
	"reflect"
	"testing"
)

func TestParseUID(t *testing.T) {
	tests := []struct {
		input    string
		expected UID
		wantErr  bool
	}{
		// Test cases without spaces in the ID
		{"AWS::Account::12345abc", UID{Type: "AWS::Account", ID: "12345abc"}, false},
		{"GCP::Project::dev", UID{Type: "GCP::Project", ID: "dev"}, false},
		{"MyNamespace::Nested::Nested::something", UID{Type: "MyNamespace::Nested::Nested", ID: "something"}, false},

		// Test case with spaces in the ID and enclosed in quotes
		{"ExampleWithSpaces::\"some ID here with spaces\"", UID{Type: "ExampleWithSpaces", ID: "some ID here with spaces"}, false},

		// Test case with spaces in the ID but not enclosed in quotes (should be invalid)
		{"ExampleWithSpaces::some ID here with spaces", UID{}, true},

		// Test case with spaces in the type (should be invalid)
		{"Type With Spaces::someID", UID{}, true},

		// Test case with invalid format (not enough parts)
		{"InvalidFormat", UID{}, true},

		// Test case with invalid format (empty input)
		{"", UID{}, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseUID(test.input)

			// Check for error
			if (err != nil) != test.wantErr {
				t.Errorf("ParseUID() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			// Compare the results
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ParseUID() got = %v, want %v", result, test.expected)
			}
		})
	}
}
