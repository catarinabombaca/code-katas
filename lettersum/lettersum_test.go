package lettersum_test

import (
	"github.com/catarinabombaca/code-katas/lettersum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterSum(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{"", 0},
			{"a", 1},
			{"z", 26},
			{"cab", 6},
			{"excellent", 100},
			{"microspectrophotometries", 317},
			{"cAb", 6},
			{"EXCELLENT", 100},
		}
	)

	for _, test := range tests {
		result := lettersum.Execute(test.input)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input      string
	expectedResult int
}