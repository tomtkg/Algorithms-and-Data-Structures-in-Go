package main

import (
	"mylib"
)

func Example_mainA() {
	mylib.SetString("*\n1.2\n+\n3.4\n1.6")
	main()
	// Output:
	// 6.00
}
func Example_mainB() {
	mylib.SetString("+\n-0.1\n*\n-\n2.5\n4.7\n+\n1.7\n2.8")
	main()
	// Output:
	// -10.00
}
