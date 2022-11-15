package main

import "mylib"

func Example_mainA() {
	data := []int{0, 23, 4, 19, 27, 2, 31, 10, 7, 9, 19}
	mylib.SetStdin(data)
	main()
	// Output:
	// ** Before heap...
	// [0 23 4 19 27 2 31 10 7 9 19]
	// Result i = 5
	// [0 23 4 19 27 19 31 10 7 9 2]
	// Result i = 4
	// [0 23 4 19 27 19 31 10 7 9 2]
	// Result i = 3
	// [0 23 4 31 27 19 19 10 7 9 2]
	// Result i = 2
	// [0 23 27 31 9 19 19 10 7 4 2]
	// Result i = 1
	// [0 31 27 23 9 19 19 10 7 4 2]
}

func Example_mainB() {
	data := []int{0, 55, 30, 97, 31, 85, 50, 41, 86, 62, 47}
	mylib.SetStdin(data)
	main()
	// Output:
	// ** Before heap...
	// [0 55 30 97 31 85 50 41 86 62 47]
	// Result i = 5
	// [0 55 30 97 31 85 50 41 86 62 47]
	// Result i = 4
	// [0 55 30 97 86 85 50 41 31 62 47]
	// Result i = 3
	// [0 55 30 97 86 85 50 41 31 62 47]
	// Result i = 2
	// [0 55 86 97 62 85 50 41 31 30 47]
	// Result i = 1
	// [0 97 86 55 62 85 50 41 31 30 47]
}
