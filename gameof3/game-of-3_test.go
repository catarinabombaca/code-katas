package gameof3_test

import (
	"github.com/catarinabombaca/code-katas/gameof3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameof3(t *testing.T) {
	var (
		assert = assert.New(t)
		num int64 = 100
		expectedResult = []int64{100, 99, 33, 11, 12, 4, 3, 1}
	)

	result := gameof3.Execute(num)

	assert.Equal(expectedResult, result)
}
