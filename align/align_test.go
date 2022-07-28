package align_test

import (
	"errors"
	"github.com/catarinabombaca/code-katas/align"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlign(t *testing.T) {
	var (
		assert = assert.New(t)
		tests  = []test{
			{"SaltPay", 11, align.L, "SaltPay....", nil},
			{"SaltPay", 11, align.R, "....SaltPay", nil},
			{"SaltPay", 11, align.C, "..SaltPay..", nil},
			{"SaltPay", 12, align.C, "...SaltPay..", nil},
			{"SaltPay", 3, align.C, "", errors.New("word cannot be bigger than length")},
			{"", 6, align.C, "......", nil},
		}
	)

	for _, test := range tests {
		result, err := align.Execute(test.word, test.length, test.direction)

		assert.Equal(test.expectedError, err)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	word           string
	length         int
	direction      align.Direction
	expectedResult string
	expectedError  error
}
