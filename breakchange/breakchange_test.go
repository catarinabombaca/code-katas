package breakchange_test

import (
	"github.com/catarinabombaca/code-katas/breakchange"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakIntoChange(t *testing.T) {
	var (
		assert = assert.New(t)
		tests  = []test{
			{3.45, map[string]int{"£2": 1, "£1": 1, "20p": 2, "5p": 1}},
			{160, map[string]int{"£50": 3, "£10": 1}},
			{0.13, map[string]int{"10p": 1, "2p": 1, "1p": 1}},
		}
	)

	for _, test := range tests {
		result := breakchange.Execute(test.amount)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	amount         float64
	expectedResult map[string]int
}
