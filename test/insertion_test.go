package test

import (
	"goAlgorithms/sorts"
	"testing"
)

func Test_InsertionOrder(t *testing.T) {
	var arrayToOrder = []int{3, 4, 5, 2, 15}
	var expected = []int{2, 3, 4, 5, 15}
	result := sorts.Insertion(arrayToOrder)
	if !compareResult(expected, result) {
		t.Error("not ordered")
	}
}

func Test_BigInsertionOrder(t *testing.T) {
	var arrayToOrder = []int{45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45}
	result := sorts.Insertion(arrayToOrder)
	if !compareResult(expected, result) {
		t.Error("not ordered")
	}
}

func Test_InsertionDescOrder(t *testing.T) {
	var arrayToOrder = []int{3, 4, 5, 2, 15}
	var expected = []int{15, 5, 4, 3, 2}
	result := sorts.InsertionDesc(arrayToOrder)
	if !compareResult(expected, result) {
		t.Error("not ordered")
	}
}

func Test_BigInsertionDescOrder(t *testing.T) {
	var arrayToOrder = []int{45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var expected = []int{45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := sorts.InsertionDesc(arrayToOrder)
	if !compareResult(expected, result) {
		t.Error("not ordered")
	}
}
