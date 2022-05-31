package fibonacci

func Execute(n int) int {
	sequence := &[]int{0,1}
	createFibonnaci(sequence, n)
	lastValue := (*sequence)[n]
	return lastValue
}

func createFibonnaci(sequence *[]int, n int) {
	shouldAddNumber := len(*sequence)-1 < n && n > 1
	if shouldAddNumber {
		secondToLastNumber := (*sequence)[len(*sequence)-2]
		lastNumber := (*sequence)[len(*sequence)-1]
		*sequence = append(*sequence, secondToLastNumber+lastNumber)
		createFibonnaci(sequence, n)
	}
}
