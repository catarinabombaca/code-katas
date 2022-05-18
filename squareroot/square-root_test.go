package squareroot_test

import (
	"github.com/catarinabombaca/code-katas/squareroot"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquareRoot(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{[]int64{4,3,9,7,2,1}, []int64{2,9,3,49,4,1}},
			{[]int64{40,36,14,478,1,1}, []int64{1600,6,196,228484,1,1}},
		}
	)

	for _, test := range tests {
		result := squareroot.Execute(test.input)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input []int64
	expectedResult []int64
}
