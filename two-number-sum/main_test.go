package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase(t *testing.T) {
	type testCase struct {
		name      string
		input     []int
		targetSum int
		expected  []int
	}
	cases := []*testCase{
		{
			name:      "Test 1",
			input:     []int{3, 5, -4, 8, 11, 1, -1, 6},
			targetSum: 10,
			expected:  []int{11, -1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := TwoNumberSum(c.input, c.targetSum)
			assert.Equal(t, c.expected, output)
		})
	}
}
