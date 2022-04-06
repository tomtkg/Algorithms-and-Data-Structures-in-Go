package main

import (
	"os"
)

func Example_mainA() {
	r, w, _ := os.Pipe()
	w.Write([]byte("*\n1.2\n+\n3.4\n1.6"))
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// 6.00
}
func Example_mainB() {
	r, w, _ := os.Pipe()
	w.Write([]byte("+\n-0.1\n*\n-\n2.5\n4.7\n+\n1.7\n2.8"))
	w.Close()
	os.Stdin = r
	main()
	// Output:
	// -10.00
}
