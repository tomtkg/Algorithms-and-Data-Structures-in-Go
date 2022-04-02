package main

import (
	"testing"
)

func Example_radixSortA() {
	radixSort([]int{374, 889, 309, 397, 987, 473, 346, 607, 881})
	// Output:
	// [374 889 309 397 987 473 346 607 881]
	// [881 473 374 346 397 987 607 889 309]
	// [607 309 346 473 374 881 987 889 397]
	// [309 346 374 397 473 607 881 889 987]
}
func Example_radixSortB() {
	radixSort([]int{523, 39, 785, 257, 351, 616, 292, 788, 802})
	// Output:
	// [523 39 785 257 351 616 292 788 802]
	// [351 292 802 523 785 616 257 788 39]
	// [802 616 523 39 351 257 785 788 292]
	// [39 257 292 351 523 616 785 788 802]
}

func Test_getSpecifiedDigitA(t *testing.T) {
	for i, want := range []int{-1, 7, 8, 0, -1} {
		if got := getSpecifiedDigit(87, i); got != want {
			t.Errorf("getSpecifiedDigit() = %v, want %v", got, want)
		}
	}
}
func Test_getSpecifiedDigitB(t *testing.T) {
	for i, want := range []int{-1, 5, 4, 3, -1} {
		if got := getSpecifiedDigit(345, i); got != want {
			t.Errorf("getSpecifiedDigit() = %v, want %v", got, want)
		}
	}
}
func Test_getSpecifiedDigitC(t *testing.T) {
	for i, want := range []int{-1, 9, 8, 7, -1} {
		if got := getSpecifiedDigit(6789, i); got != want {
			t.Errorf("getSpecifiedDigit() = %v, want %v", got, want)
		}
	}
}
func Test_getSpecifiedDigitD(t *testing.T) {
	for i, want := range []int{-1, 9, 8, 7, -1} {
		if got := getSpecifiedDigit(-6789, i); got != want {
			t.Errorf("getSpecifiedDigit() = %v, want %v", got, want)
		}
	}
}
