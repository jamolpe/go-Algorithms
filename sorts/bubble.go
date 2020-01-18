package sorts

// Bubble : ASC order
func Bubble(slice []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				value := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = value
				sorted = false
			}
		}
	}
	return slice
}

// BubbleDesc : desc order
func BubbleDesc(slice []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] < slice[i+1] {
				value := slice[i]
				slice[i] = slice[i+1]
				slice[i+1] = value
				sorted = false
			}
		}
	}
	return slice
}
