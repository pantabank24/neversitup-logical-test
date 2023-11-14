package unittest

import (
	"fmt"
	"testing"

	"github.com/pantabank24/logical-test/project"
	"github.com/stretchr/testify/assert"
)

type testCountSmiley struct {
	input  []string
	expect int
}

func TestCountSmiley(t *testing.T) {
	tests := []testCountSmiley{
		{
			input:  []string{":)", ";(", ";}", ":-D"},
			expect: 2,
		},
		{
			input:  []string{";D", ":-(", ":-)", ";~)"},
			expect: 3,
		},
		{
			input:  []string{";]", ":[", ";*", ":$", ";-D"},
			expect: 1,
		},
		{
			input:  []string{},
			expect: 0,
		},
	}

	for _, test := range tests {

		result := project.CountSmileys(test.input)

		fmt.Println("should return", test.expect)

		assert.Equal(t, test.expect, result)
	}
}
