package lettersum

import (
	"strings"
)

func Execute(input string) int {
	lettersAsValues := []rune(strings.ToLower(input))
	sum := sumAllValues(lettersAsValues)

	return sum
}

func sumAllValues(lettersAsValues []rune) int {
	sum := 0
	for _, value := range lettersAsValues {
		sum += int(value) - 96
	}
	return sum
}