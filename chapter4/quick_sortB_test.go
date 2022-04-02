package main

import "testing"

func Test_quickSortB1(t *testing.T) {
	seq := []int{37, 33, 29, 25, 22, 20, 19, 14, 10, 8}
	want := []int{8, 10, 14, 19, 20, 22, 25, 29, 33, 37}
	quickSortB(seq, 0, len(seq)-1)
	for i := range seq {
		if seq[i] != want[i] {
			t.Errorf("seq[%d] = %v, want %v", i, seq[i], want[i])
		}
	}
}
func Test_quickSortB2(t *testing.T) {
	seq := []int{6, 9, 12, 7, 15, 23, 2, 10, 4, 20}
	want := []int{2, 4, 6, 7, 9, 10, 12, 15, 20, 23}
	quickSortB(seq, 0, len(seq)-1)
	for i := range seq {
		if seq[i] != want[i] {
			t.Errorf("seq[%d] = %v, want %v", i, seq[i], want[i])
		}
	}
}
func Test_quickSortB3(t *testing.T) {
	seq := []int{6, 9, 3, 8, 7, 5, 4, 2, 1}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	quickSortB(seq, 0, len(seq)-1)
	for i := range seq {
		if seq[i] != want[i] {
			t.Errorf("seq[%d] = %v, want %v", i, seq[i], want[i])
		}
	}
}
