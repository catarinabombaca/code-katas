package rotate_test

import (
	"errors"
	"github.com/catarinabombaca/code-katas/rotate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{[]int64{1,2,3,4,5,6}, 2, []int64{3,4,5,6,1,2}, nil},
			{[]int64{1,2,3,4,5,6}, 6, []int64{1,2,3,4,5,6}, nil},
			{[]int64{1,2,3,4,5,6}, 7, nil, errors.New("rotation value is not allowed")},
			{[]int64{1,2,3,4,5,6}, 0, []int64{1,2,3,4,5,6}, nil},
			{[]int64{1,2,3,4,5,6}, -1, nil, errors.New("rotation value is not allowed")},
		}
	)

	for _, test := range tests {
		result, err := rotate.Execute(test.input, test.rotation)
		if err != nil {
			assert.EqualError(err, test.expectedError.Error())
		}
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input []int64
	rotation int
	expectedResult []int64
	expectedError error
}
