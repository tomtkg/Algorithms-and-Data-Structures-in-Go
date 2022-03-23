package main

import (
	"os"
)

func Example_main() {
	for _, v := range []string{"a", "b"} {
		os.Stdin, _ = os.Open(v + ".txt")
		main()
	}
	// Output:
	// 6.00
	// -10.00
}
