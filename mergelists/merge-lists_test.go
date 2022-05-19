package mergelists_test

import (
	"github.com/catarinabombaca/code-katas/mergelists"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeLists(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{[]int{1, 4, 6}, []int{2, 3, 5}, []int{1, 2, 3, 4, 5, 6}},
			{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		}
	)

	for _, test := range tests {
		result := mergelists.Execute(test.firstList, test.secondList)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	firstList  []int
	secondList []int
	expectedResult []int
}
