package main

import (
	"os"
)

func ExampleMain() {
	main()
	// Output:
}

func ExampleMainA() {
	os.Args[1] = "a"
	main()
	// Output:
	// 55
}

func ExampleMainB() {
	os.Args[1] = "b"
	main()
	// Output:
	// 55
}

func ExampleMainC() {
	os.Args[1] = "c"
	main()
	// Output:
}
