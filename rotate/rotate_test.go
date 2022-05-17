package rotate_test

import (
	"github.com/catarinabombaca/code-katas/rotate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{[]int64{1,2,3,4,5,6}, 2, []int64{3,4,5,6,1,2}},
		}
	)

	for _, test := range tests {
		result := rotate.Execute(test.input, test.rotation)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input []int64
	rotation int64
	expectedResult []int64
}
