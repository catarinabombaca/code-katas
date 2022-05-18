package squareroot

import (
	"math"
)

func Execute(input []int64) []int64 {
	for i, value := range input {
		squareRootFromValue := math.Sqrt(float64(value))

		if squareRootFromValue - math.Floor(squareRootFromValue) == 0 {
			input[i] = int64(squareRootFromValue)
		} else {
			input[i] = value*value
		}
	}

	return input
}