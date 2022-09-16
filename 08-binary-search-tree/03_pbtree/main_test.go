package main

func Example_a() {
	stack, sp = []int{}, 0
	a()
	// Output:
	// 図8.8
	// -- Binary tree --
	// 0
	// 	10
	// 		20
	// 			30
	// 				40
	//
	// -- inorder trversal --
	// 0
	// 10
	// 20
	// 30
	// 40
	//
	// -- Generated binary tree --
	// 	0
	// 		10
	// 20
	// 	30
	// 		40
}

func Example_b() {
	stack, sp = []int{}, 0
	b()
	// Output:
	// 図8.9
	// -- Generated binary tree --
	// 		11
	// 	12
	// 		13
	// 			14
	// 15
	// 		16
	// 			17
	// 	18
	// 		19
	// 			20
}

func Example_c() {
	stack, sp = []int{}, 0
	c()
	// Output:
	// 章末(1)
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
	// (a)
	//  -> 42 -> 76 -> 63
	// (b)
	// -- preorder trversal --
	//  -> 42 -> 15 -> 18 -> 28 -> 76 -> 63 -> 98 -> 92
	// -- postorder trversal --
	//  -> 28 -> 18 -> 15 -> 63 -> 92 -> 98 -> 76 -> 42
	// (c)
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

func Example_d() {
	stack, sp = []int{}, 0
	d()
	// Output:
	// 章末(2)
	// -- Generated binary tree --
	//	2
	// 		9
	// 11
	// 		28
	// 	31
	// 		45
}
