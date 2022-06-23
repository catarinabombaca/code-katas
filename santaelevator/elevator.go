package santaelevator

func Execute(instructions string) int {
	var floor = 0

	for _, instruction := range instructions {
		if instruction == 40 {
			floor++
		}
		if instruction == 41 {
			floor--
		}
	}

	return floor
}
