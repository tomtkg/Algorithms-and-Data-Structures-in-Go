package main

func Example_mainA() {
	main()
	// Output:
	// -- Binary tree --
	// 	15
	// 		18
	// 			28
	// 42
	// 		63
	// 	76
	// 			92
	// 		98
	//
	// (a) 63: -> 42 -> 76 -> 63
	// -- preorder trversal --
	//  -> 42 -> 15 -> 18 -> 28 -> 76 -> 63 -> 98 -> 92
	// -- postorder trversal --
	//  -> 28 -> 18 -> 15 -> 63 -> 92 -> 98 -> 76 -> 42
	// -- Generated binary tree --
	// 	15
	// 		18
	// 			28
	// 				30
	// 42
	// 		63
	// 	76
	// 			92
	// 		98
}
