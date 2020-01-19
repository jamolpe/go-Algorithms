package sorts

func orderSides(first []int, second []int) []int {
	i, j := 0, 0
	result := make([]int, len(first)+len(second))

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			result[i+j] = first[i]
			i++
		} else {
			result[i+j] = second[j]
			j++
		}
	}

	// WE COMPLETE THE RESULT WITH THE REST THAT HAS NO EQUIVALENT POSITION IN THE OTHER
	for i < len(first) {
		result[i+j] = first[i]
		i++
	}

	for j < len(second) {
		result[i+j] = second[j]
		j++
	}

	return result
}

// MergeSort : order function to merge
func MergeSort(sliceToOrder []int) []int {
	if len(sliceToOrder) < 2 {
		return sliceToOrder
	}

	middleIndex := len(sliceToOrder) / 2
	firstMergeSide := MergeSort(sliceToOrder[:middleIndex])
	secondMergeSide := MergeSort(sliceToOrder[middleIndex:])
	return orderSides(firstMergeSide, secondMergeSide)
}
