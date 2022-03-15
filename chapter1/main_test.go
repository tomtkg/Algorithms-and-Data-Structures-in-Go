package main

import (
	"os"
)

func Example_main() {
	main()
	// Output:
}

func Example_mainA() {
	os.Args[1] = "a"
	main()
	// Output:
	// 55
}

func Example_mainB() {
	os.Args[1] = "b"
	main()
	// Output:
	// 55
}

func Example_mainC() {
	os.Args[1] = "c"
	main()
	// Output:
}
