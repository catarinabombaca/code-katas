package rotate

import (
	"errors"
)

func Execute(input []int64, rotation int) ([]int64, error) {
	if rotation > len(input) || rotation < 0 {
		return nil, errors.New("rotation value is not allowed")
	}

	return append(input[rotation:], input[:rotation]...), nil
}
