package sorts

import "sync"

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

// Merge : using chanels to calculate it NOT OPTIMAL APROACH USING A SEMAPHORE SHOULD BE BETTER
func Merge(sliceToOrder []int, c chan []int) {
	if len(sliceToOrder) < 2 {
		c <- sliceToOrder
	}
	middleIndex := len(sliceToOrder) / 2
	leftChannel := make(chan []int)
	rigthChannel := make(chan []int)
	go Merge(sliceToOrder[:middleIndex], leftChannel)
	go Merge(sliceToOrder[middleIndex:], rigthChannel)
	firstMergeSide, secondMergeSide := <-leftChannel, <-rigthChannel
	c <- orderSides(firstMergeSide, secondMergeSide)
}

// MergeWithSemaphore : using semaphore to order
func MergeWithSemaphore(sliceToOrder []int, c chan struct{}) []int {
	if len(sliceToOrder) < 2 {
		return sliceToOrder
	}
	middleIndex := len(sliceToOrder) / 2

	// wait max 2 channels
	wg := sync.WaitGroup{}
	wg.Add(2)

	var lMerge []int
	var rMerge []int

	select {
	case c <- struct{}{}:
		go func() {
			lMerge = MergeWithSemaphore(sliceToOrder[:middleIndex], c)
			<-c
			wg.Done()
		}()
	default:
		lMerge = MergeSort(sliceToOrder[:middleIndex])
		wg.Done()
	}
	select {
	case c <- struct{}{}:
		go func() {
			rMerge = MergeWithSemaphore(sliceToOrder[middleIndex:], c)
			<-c
			wg.Done()
		}()
	default:
		rMerge = MergeSort(sliceToOrder[middleIndex:])
		wg.Done()
	}
	wg.Wait()
	return orderSides(lMerge, rMerge)
}
