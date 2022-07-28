package align

import (
	"errors"
	"math"
	"strings"
)

const L Direction = "L"
const C Direction = "C"
const R Direction = "R"

type Direction string

func Execute(word string, length int, direction Direction) (string, error) {
	if len(word) > length {
		return "", errors.New("word cannot be bigger than length")
	}

	var (
		finalStringLength = length - len(word)
		sb                strings.Builder
	)

	switch direction {
	case L:
		sb.WriteString(word)
		addDots(&sb, finalStringLength)
	case R:
		addDots(&sb, finalStringLength)
		sb.WriteString(word)
	case C:
		leftDotsLength := int(math.Round(float64(finalStringLength) / 2))
		rightDotsLength := int(math.Floor(float64(finalStringLength) / 2))
		addDots(&sb, leftDotsLength)
		sb.WriteString(word)
		addDots(&sb, rightDotsLength)
	}

	return sb.String(), nil
}

func addDots(builder *strings.Builder, length int) {
	for i := 0; i < length; i++ {
		builder.WriteString(".")
	}
}
