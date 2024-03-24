package ranker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexStringComparison(t *testing.T) {
	tests := []struct {
		name                        string
		leftHandSide, rightHandSide string
	}{
		{
			name:          "single character: letter order",
			leftHandSide:  "f",
			rightHandSide: "a",
		},
		{
			name:          "single character: letter/digit order",
			leftHandSide:  "a",
			rightHandSide: "9",
		},
		{
			name:          "single character: digit order",
			leftHandSide:  "7",
			rightHandSide: "3",
		},
		{
			name:          "multiple letters",
			leftHandSide:  "abc",
			rightHandSide: "aab",
		},
		{
			name:          "multiple digits",
			leftHandSide:  "781",
			rightHandSide: "703",
		},
		{
			name:          "mixed characters",
			leftHandSide:  "d81b7a",
			rightHandSide: "a81b7a",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.True(t, testCase.leftHandSide > testCase.rightHandSide)
		})
	}
}
