package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseFormDate(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "ValidDate",
			input:    "2023-12-23",
			expected: time.Date(2023, time.December, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "InvalidDate",
			input:       "2023-13-40",
			expected:    time.Time{},
			expectError: true,
		},
	}

	// Execute each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseFormDate(tc.input)
			if tc.expectError {
				assert.Error(t, err, "Expected an error for input: "+tc.input)
			} else {
				assert.NoError(t, err, "Did not expect an error for input: "+tc.input)
				assert.Equal(t, tc.expected, result, "Expected and actual time do not match for input: "+tc.input)
			}
		})
	}
}
