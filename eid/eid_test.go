package eid

import (
	"reflect"
	"testing"
)

func TestParseEID(t *testing.T) {
	tests := []struct {
		input    string
		expected EID
		parser   *Parser
		wantErr  bool
	}{
		// Test cases without spaces in the ID
		{input: "AWS::Account::12345abc", expected: EID{Type: "AWS::Account", ID: "12345abc"}, wantErr: false},
		{input: "GCP::Project::dev", expected: EID{Type: "GCP::Project", ID: "dev"}, wantErr: false},
		{input: "MyNamespace::Nested::Nested::something", expected: EID{Type: "MyNamespace::Nested::Nested", ID: "something"}, wantErr: false},

		// Test case with spaces in the ID and enclosed in quotes
		{input: "ExampleWithSpaces::\"some ID here with spaces\"", expected: EID{Type: "ExampleWithSpaces", ID: "some ID here with spaces"}, wantErr: false},

		// Test case with spaces in the ID but not enclosed in quotes (should be invalid)
		{input: "ExampleWithSpaces::some ID here with spaces", expected: EID{}, wantErr: true},

		// Test case with spaces in the type (should be invalid)
		{input: "Type With Spaces::someID", expected: EID{}, wantErr: true},

		// Test case with invalid format (not enough parts)
		{input: "InvalidFormat", expected: EID{}, wantErr: true},

		// Test case with invalid format (empty input)
		{input: "", expected: EID{}, wantErr: true},

		// custom parser correctly identifies that a lookup with colons is not an eid because a single colon preceeds a double
		{input: "arn:aws:sso:::permissionSet/ssoins-1234eeee/ps-1234eee", expected: EID{}, wantErr: true},

		// don't split triple colons or more
		{input: "Before::After:::Triple", expected: EID{Type: "Before", ID: "After:::Triple"}, wantErr: false},
		// don't split triple+ colons
		{input: "Before::After::::Triple", expected: EID{Type: "Before", ID: "After::::Triple"}, wantErr: false},
		// don't split triple+ colons
		{input: "Before:Hello::After::::Triple", expected: EID{}, wantErr: true},

		{input: "AWS::Account::12345abc", expected: EID{}, wantErr: true, parser: &Parser{RequireQuotedID: true}},
		{input: `AWS::Account::"12345abc"`, expected: EID{Type: "AWS::Account", ID: "12345abc"}, wantErr: false, parser: &Parser{RequireQuotedID: true}},
		{input: `Access::LinkedIdentity::"AWS::IDC::User::1234"`, expected: EID{Type: "Access::LinkedIdentity", ID: "AWS::IDC::User::1234"}, wantErr: false, parser: &Parser{RequireQuotedID: true}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			var result EID
			var err error
			if test.parser != nil {
				result, err = test.parser.Parse(test.input)
			} else {
				result, err = Parse(test.input)
			}

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
