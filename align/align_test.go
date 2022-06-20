package align_test

import (
	"github.com/catarinabombaca/code-katas/align"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlign(t *testing.T) {
	var (
		assert = assert.New(t)
		tests = []test{
			{"SaltPay", 11, "L", "SaltPay...."},
			{"SaltPay", 11, "R", "....SaltPay"},
			{"SaltPay", 11, "C", "..SaltPay.."},
			{"SaltPay", 12, "C", "...SaltPay.."},
		}
	)

	for _, test := range tests {
		result := align.Execute(test.word, test.length, test.direction)
		assert.Equal(test.expectedResult, result)
	}
}

type test struct {
	word string
	length int
	direction string
	expectedResult string
}