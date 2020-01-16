package test

import (
	"goAlgorithms/sorts"
	"testing"
)

func Test_BubbleOrder(t *testing.T) {
	var arrayToOrder = []int{3, 4, 5, 2, 15}
	var expected = []int{2, 3, 4, 5, 15}
	result := sorts.Bubble(arrayToOrder)
	if !compareResult(expected, result) {
		t.Error("not ordered")
	}
}

func compareResult(gived []int, expected []int) bool {
	if len(gived) != len(expected) {
		return false
	}
	for pos := range gived {
		if gived[pos] != expected[pos] {
			return false
		}
	}
	return true
}
