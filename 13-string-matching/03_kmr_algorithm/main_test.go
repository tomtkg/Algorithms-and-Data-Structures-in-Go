package main

import "mylib"

func Example_main() {
	mylib.SetStdin([][2]string{
		{"testdtested", "tested"},
		{"testested", "tested"},
		{"testested", "go"},
		{"algorithm", "go"},
	})
	main()
	// Output:
	// text = testdtested
	// word = tested
	// 5
	// text = testested
	// word = tested
	// 3
	// text = testested
	// word = go
	// 9
	// text = algorithm
	// word = go
	// 2
}
