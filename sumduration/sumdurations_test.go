package sumduration_test

import (
	"github.com/catarinabombaca/code-katas/sumduration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumDuration(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{[]string{"12:32", "34:01", "15:23", "9:27", "55:22", "25:56"}, "2:32:41"},
			{[]string{"12:32", "34:01", "15:23", "9:27", "25:22", "25:56"}, "2:02:41"},
		}
	)

	for _, test := range tests {
		result := sumduration.Execute(test.input)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input          []string
	expectedResult string
}

