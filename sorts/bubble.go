package sorts

func Bubble(array []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				value := array[i]
				array[i] = array[i+1]
				array[i+1] = value
				sorted = false
			}
		}
	}
	return array
}
