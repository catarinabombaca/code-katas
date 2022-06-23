package santaelevator

func Execute(instructions string) (int, int) {
	var currentFloor = 0
	var position = 0

	for i, instruction := range instructions {
		if instruction == 40 {
			currentFloor++
		}

		if instruction == 41 {
			currentFloor--
		}

		if isBasementPosition(currentFloor) && position == 0 {
			position += i+1
		}
	}

	return currentFloor, position
}

func isBasementPosition(floor int) bool {
	if floor == -1 {
		return true
	} else {
		return false
	}
}
