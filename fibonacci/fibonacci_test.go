package fibonacci_test

import (
	"github.com/catarinabombaca/code-katas/fibonacci"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{0, 0},
			{4, 3},
			{15, 610},
			{92, 7540113804746346429},
		}
	)

	for _, test := range tests {
		result := fibonacci.Execute(test.input)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	input      int
	expectedResult int
}