package main

import (
	"fmt"
	"testing"
)

func Example_radixSort() {
	radixSort([]int{374, 889, 309, 397, 987, 473, 346, 607, 881})
	fmt.Println()
	radixSort([]int{523, 39, 785, 257, 351, 616, 292, 788, 802})
	// Output:
	// [374 889 309 397 987 473 346 607 881]
	// [881 473 374 346 397 987 607 889 309]
	// [607 309 346 473 374 881 987 889 397]
	// [309 346 374 397 473 607 881 889 987]
	//
	// [523 39 785 257 351 616 292 788 802]
	// [351 292 802 523 785 616 257 788 39]
	// [802 616 523 39 351 257 785 788 292]
	// [39 257 292 351 523 616 785 788 802]
}

func Test_getSpecifiedDigit(t *testing.T) {
	tests := []struct {
		name string
		num  int
		m    map[int]int
	}{
		{"2桁", 87, map[int]int{0: -1, 1: 7, 2: 8, 3: 0, 4: -1}},
		{"3桁", 345, map[int]int{0: -1, 1: 5, 2: 4, 3: 3, 4: -1}},
		{"4桁", 6789, map[int]int{0: -1, 1: 9, 2: 8, 3: 7, 4: -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, want := range tt.m {
				if got := getSpecifiedDigit(tt.num, i); got != want {
					t.Errorf("getSpecifiedDigit() = %v, want %v", got, want)
				}
			}
		})
	}
}
