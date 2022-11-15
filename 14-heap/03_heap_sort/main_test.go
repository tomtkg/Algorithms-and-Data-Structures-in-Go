package main

import "mylib"

func Example_mainA() {
	data := []int{0, 23, 4, 19, 27, 2, 31, 10, 7, 9, 19}
	mylib.SetStdin(data)
	main()
	// Output:
	// ** Before heap...
	// [0 23 4 19 27 2 31 10 7 9 19]
	// ** After heap...
	// [0 31 27 23 9 19 19 10 7 4 2]
	// ** After sort...
	// [0 2 4 7 9 10 19 19 23 27 31]
}

func Example_mainB() {
	data := []int{0, 7, 27, 19, 9, 10, 23, 2, 12, 4, 31}
	mylib.SetStdin(data)
	main()
	// Output:
	// ** Before heap...
	// [0 7 27 19 9 10 23 2 12 4 31]
	// ** After heap...
	// [0 31 27 23 12 10 19 2 9 4 7]
	// ** After sort...
	// [0 2 4 7 9 10 12 19 23 27 31]
}

func Example_mainC() {
	data := []int{0, 26, 16, 32, 15, 40, 13, 23, 33, 14, 12}
	mylib.SetStdin(data)
	main()
	// Output:
	// ** Before heap...
	// [0 26 16 32 15 40 13 23 33 14 12]
	// ** After heap...
	// [0 40 33 32 26 16 13 23 15 14 12]
	// ** After sort...
	// [0 12 13 14 15 16 23 26 32 33 40]
}
