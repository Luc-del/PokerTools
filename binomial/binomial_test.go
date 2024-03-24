package binomial

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoefficient(t *testing.T) {
	tests := []struct {
		setSize    int
		subsetSize int
		expected   int
	}{
		{
			setSize:    1,
			subsetSize: 1,
			expected:   1,
		},
		{
			setSize:    3,
			subsetSize: 1,
			expected:   3,
		},
		{
			setSize:    3,
			subsetSize: 2,
			expected:   3,
		},
		{
			setSize:    5,
			subsetSize: 2,
			expected:   10,
		},
		{
			setSize:    7,
			subsetSize: 2,
			expected:   21,
		},
		{
			setSize:    7,
			subsetSize: 5,
			expected:   21,
		},
		{
			setSize:    8,
			subsetSize: 4,
			expected:   70,
		},
		{
			setSize:    48,
			subsetSize: 5,
			expected:   1712304,
		},
		{
			setSize:    52,
			subsetSize: 2,
			expected:   1326,
		},
		{
			setSize:    52,
			subsetSize: 5,
			expected:   2598960,
		},
		{
			setSize:    52,
			subsetSize: 7,
			expected:   133784560,
		},
		{
			setSize:    64,
			subsetSize: 17,
			expected:   1379370175283520,
		},
	}
	for _, testCase := range tests {
		t.Run(fmt.Sprintf("%d within %d", testCase.subsetSize, testCase.setSize), func(t *testing.T) {
			assert.Equal(t, testCase.expected, Coefficient(testCase.setSize, testCase.subsetSize))
		})
	}
}
