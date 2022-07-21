package breakchange

import (
	"fmt"
	"math"
)

func Execute(amount float64) map[string]int {
	var (
		cashAvailable = []int{50, 20, 10, 5, 2, 1}
		cashBreakdown = make(map[string]int)
	)
	notesValue := math.Floor(amount)
	for _, value := range cashAvailable {
		if value > int(notesValue) {
			continue
		}

		cashBreakdown[fmt.Sprintf("£%d", value)] = int(notesValue) / value
	}
	return cashBreakdown
}
