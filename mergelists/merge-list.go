package mergelists

func Execute(firstList []int, secondList []int) []int {
	var (
		sortedList  []int
		firstIndex  = 0
		secondIndex = 0
	)

	for {
		if firstIndex == len(firstList)  {
			sortedList = append(sortedList, secondList[secondIndex:]...)
			break
		}

		if secondIndex == len(secondList) {
			sortedList = append(sortedList, firstList[firstIndex:]...)
			break
		}

		if firstList[firstIndex] < secondList[secondIndex] {
			sortedList = append(sortedList, firstList[firstIndex])
			firstIndex++
		} else {
			sortedList = append(sortedList, secondList[secondIndex])
			secondIndex++
		}
	}

	return sortedList
}
