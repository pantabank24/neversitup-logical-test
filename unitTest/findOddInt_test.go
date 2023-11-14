package unittest

import (
	"fmt"
	"testing"

	"github.com/pantabank24/logical-test/project"
	"github.com/stretchr/testify/assert"
)

type testFindOddInt struct {
	input  []int
	expect int
}

func TestFindOddInt(t *testing.T) {
	tests := []testFindOddInt{
		{
			input:  []int{7},
			expect: 7,
		},
		{
			input:  []int{0},
			expect: 0,
		},
		{
			input:  []int{1, 1, 2},
			expect: 2,
		},
		{
			input:  []int{0, 1, 0, 1, 0},
			expect: 0,
		},
		{
			input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expect: 4,
		},
	}

	for _, test := range tests {
		result := project.FindOddInt(test.input)

		fmt.Println(test.input, "should return", result)

		assert.Equal(t, test.expect, result)
	}
}
