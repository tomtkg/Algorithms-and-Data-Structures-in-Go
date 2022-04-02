package main

import (
	"os"
)

func Example_mainA() {
	r, w, _ := os.Pipe()
	w.Write([]byte("[37,33,29,25,22,20,19,14,10,8]"))
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// [37 33 29 25 22 20 19 14 10 8]
	// [8 10 14 19 20 22 25 29 33 37]
	// [37 33 29 25 22 20 19 14 10 8]
	// [8 10 14 19 20 22 25 29 33 37]
}
func Example_mainB() {
	r, w, _ := os.Pipe()
	w.Write([]byte("[6,9,12,7,15,23,2,10,4,20]"))
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// [6 9 12 7 15 23 2 10 4 20]
	// [2 4 6 7 9 10 12 15 20 23]
	// [6 9 12 7 15 23 2 10 4 20]
	// [2 4 6 7 9 10 12 15 20 23]
}
func Example_mainC() {
	r, w, _ := os.Pipe()
	w.Write([]byte("[6,9,3,8,7,5,4,2,1]"))
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// [6 9 3 8 7 5 4 2 1]
	// [1 2 3 4 5 6 7 8 9]
	// [6 9 3 8 7 5 4 2 1]
	// [1 2 3 4 5 6 7 8 9]
}
