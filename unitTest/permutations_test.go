package unittest

import (
	"fmt"
	"testing"

	"github.com/pantabank24/logical-test/project"
	"github.com/stretchr/testify/assert"
)

type testPermute struct {
	input  string
	expect []string
}

func TestPermute(t *testing.T) {
	tests := []testPermute{
		{
			input:  "a",
			expect: []string{"a"},
		},
		{
			input:  "ab",
			expect: []string{"ab", "ba"},
		},
		{
			input:  "abc",
			expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}

	for _, test := range tests {

		result := project.Permute("", test.input)

		fmt.Println("With input", test.input, ": ")
		fmt.Println("Your function should return", result)

		assert.Equal(t, test.expect, result)
	}
}
