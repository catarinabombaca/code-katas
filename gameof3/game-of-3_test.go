package gameof3_test

import (
	"github.com/catarinabombaca/code-katas/gameof3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameof3(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{100, []int64{100, 99, 33, 11, 12, 4, 3, 1}},
			{1, []int64{1}},
			{0, []int64{0,1}},
			{-1, []int64{-1,0,1}},
			{-40, []int64{-40, -39, -13, -12, -4, -3, -1, 0, 1}},
		}
	)

	for _, test := range tests {
	result := gameof3.Execute(test.num)
	assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	num int64
	expectedResult []int64
}
