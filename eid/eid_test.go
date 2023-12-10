package eid

import (
	"reflect"
	"testing"
)

func TestParseEID(t *testing.T) {
	tests := []struct {
		input    string
		expected EID
		wantErr  bool
	}{
		// Test cases without spaces in the ID
		{"AWS::Account::12345abc", EID{Type: "AWS::Account", ID: "12345abc"}, false},
		{"GCP::Project::dev", EID{Type: "GCP::Project", ID: "dev"}, false},
		{"MyNamespace::Nested::Nested::something", EID{Type: "MyNamespace::Nested::Nested", ID: "something"}, false},

		// Test case with spaces in the ID and enclosed in quotes
		{"ExampleWithSpaces::\"some ID here with spaces\"", EID{Type: "ExampleWithSpaces", ID: "some ID here with spaces"}, false},

		// Test case with spaces in the ID but not enclosed in quotes (should be invalid)
		{"ExampleWithSpaces::some ID here with spaces", EID{}, true},

		// Test case with spaces in the type (should be invalid)
		{"Type With Spaces::someID", EID{}, true},

		// Test case with invalid format (not enough parts)
		{"InvalidFormat", EID{}, true},

		// Test case with invalid format (empty input)
		{"", EID{}, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := Parse(test.input)

			// Check for error
			if (err != nil) != test.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			// Compare the results
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Parse() got = %v, want %v", result, test.expected)
			}
		})
	}
}
