package main

import "fmt"

func Example_selectionSort() {
	selectionSort([]int{4, 2, 3, 1}) //2line
	fmt.Println()
	selectionSort([]int{2, 3, 4, 1}) //4line
	fmt.Println()
	selectionSort([]int{1, 2, 4, 3}) //2line
	// Output:
	// [4 2 3 1]
	// [1 2 3 4]
	//
	// [2 3 4 1]
	// [1 3 4 2]
	// [1 2 4 3]
	// [1 2 3 4]
	//
	// [1 2 4 3]
	// [1 2 3 4]
}

func Example_insertionSort() {
	insertionSort([]int{4, 2, 3, 1}) //4line
	fmt.Println()
	insertionSort([]int{2, 3, 4, 1}) //2line
	fmt.Println()
	insertionSort([]int{1, 2, 4, 3}) //2line
	// Output:
	// [4 2 3 1]
	// [2 4 3 1]
	// [2 3 4 1]
	// [1 2 3 4]
	//
	// [2 3 4 1]
	// [1 2 3 4]
	//
	// [1 2 4 3]
	// [1 2 3 4]
}

func Example_bubbleSort() {
	bubbleSort([]int{4, 2, 3, 1}) //6line
	fmt.Println()
	bubbleSort([]int{2, 3, 4, 1}) //4line
	fmt.Println()
	bubbleSort([]int{1, 2, 4, 3}) //2line
	// Output:
	// [4 2 3 1]
	// [4 2 1 3]
	// [4 1 2 3]
	// [1 4 2 3]
	// [1 2 4 3]
	// [1 2 3 4]
	//
	// [2 3 4 1]
	// [2 3 1 4]
	// [2 1 3 4]
	// [1 2 3 4]
	//
	// [1 2 4 3]
	// [1 2 3 4]
}
