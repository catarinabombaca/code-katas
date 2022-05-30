package lettersum

import "strings"

func Execute(input string) int {
	runes := []rune(strings.ToLower(input))

	sum := 0
	for _, r := range runes {
		sum += int(r)-96
	}

	return sum
}
