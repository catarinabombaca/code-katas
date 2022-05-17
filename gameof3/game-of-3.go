package gameof3

func Execute(number int64) []int64 {
	var result = []int64{number}

	for number != 1 {
		switch {
		case number == 0:
			number = number+1
		case number % 3 == 0:
			number = number/3
		case (number - 1) % 3 == 0:
			number = number-1
		case (number + 1) % 3 == 0:
			number = number+1
		default:
			number = number+1
		}

		result = append(result, number)
	}

	return result
}
