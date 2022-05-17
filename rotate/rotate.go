package rotate

func Execute(input []int64, rotation int64) []int64 {
	return append(input[2:], input[:rotation]...)
}
