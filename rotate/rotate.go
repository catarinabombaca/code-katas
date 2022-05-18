package rotate

import (
	"errors"
)

func Execute(input []int64, rotation int) ([]int64, error) {
	if rotation > len(input) || rotation < 0 {
		return nil, errors.New("rotation value is not allowed")
	}

	for i := 0; i <rotation ; i++ {
		last := input[0]
		for i, _ := range input {
			if i == len(input)-1 {
				input[len(input)-1] = last
			} else {
				input[i] = input[i+1]
			}
		}
	}

	return input, nil
}