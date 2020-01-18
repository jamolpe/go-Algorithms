package sorts

// Insertion : order using insertion
func Insertion(sliceToOrder []int) []int {
	for i := 1; i < len(sliceToOrder); i++ {
		val := sliceToOrder[i]
		pos := i

		for ; pos > 0 && sliceToOrder[pos-1] >= val; pos-- {
			sliceToOrder[pos] = sliceToOrder[pos-1]
		}
		sliceToOrder[pos] = val
	}
	return sliceToOrder
}

// InsertionDesc : order descending using insertion
func InsertionDesc(sliceToOrder []int) []int {
	for i := 1; i < len(sliceToOrder); i++ {
		val := sliceToOrder[i]
		pos := i

		for ; pos > 0 && sliceToOrder[pos-1] <= val; pos-- {
			sliceToOrder[pos] = sliceToOrder[pos-1]
		}
		sliceToOrder[pos] = val
	}
	return sliceToOrder
}
