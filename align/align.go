package align

import (
	"math"
	"strings"
)

func Execute(word string, length int, direction string) string {
	var (
		finalStringLength = length - len(word)
		sb strings.Builder
	)

	switch direction {
	case "L":
		sb.WriteString(word)
		addDots(&sb, finalStringLength)
	case "R":
		addDots(&sb, finalStringLength)
		sb.WriteString(word)
	case "C":
		leftDotsLength := int(math.Round(float64(finalStringLength)/2))
		rightDotsLength := int(math.Floor(float64(finalStringLength)/2))
		addDots(&sb, leftDotsLength)
		sb.WriteString(word)
		addDots(&sb, rightDotsLength)
	}

	return sb.String()
}

func addDots(builder *strings.Builder, length int) {
	for i := 0; i < length; i++ {
		builder.WriteString(".")
	}
}
