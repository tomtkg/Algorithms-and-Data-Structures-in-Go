package main

import (
	"encoding/json"
	"os"
)

func Example_mainA() {
	graph := [][2]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{3, 6},
		{3, 7},
		{4, 5},
		{5, 7},
		{5, 8},
		{6, 7},
		{7, 8},
	}
	s, _ := json.Marshal(graph)
	r, w, _ := os.Pipe()
	w.Write(s)
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// [0 1 0 1 0 0 0 0 0]
	// [1 0 1 1 1 0 0 0 0]
	// [0 1 0 0 0 1 0 0 0]
	// [1 1 0 0 0 0 1 1 0]
	// [0 1 0 0 0 1 0 0 0]
	// [0 0 1 0 1 0 0 1 1]
	// [0 0 0 1 0 0 0 1 0]
	// [0 0 0 1 0 1 1 0 1]
	// [0 0 0 0 0 1 0 1 0]
	// (0, 1)(1, 2)(2, 5)(5, 4)(5, 7)(7, 3)(3, 6)(7, 8)
}
func Example_mainB() {
	graph := [][2]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 4},
		{2, 5},
		{3, 4},
		{3, 6},
		{4, 5},
		{4, 7},
		{5, 8},
		{6, 7},
		{7, 8},
	}
	s, _ := json.Marshal(graph)
	r, w, _ := os.Pipe()
	w.Write(s)
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// [0 1 0 1 0 0 0 0 0]
	// [1 0 1 0 1 0 0 0 0]
	// [0 1 0 0 0 1 0 0 0]
	// [1 0 0 0 1 0 1 0 0]
	// [0 1 0 1 0 1 0 1 0]
	// [0 0 1 0 1 0 0 0 1]
	// [0 0 0 1 0 0 0 1 0]
	// [0 0 0 0 1 0 1 0 1]
	// [0 0 0 0 0 1 0 1 0]
	// (0, 1)(1, 2)(2, 5)(5, 4)(4, 3)(3, 6)(6, 7)(7, 8)
}
