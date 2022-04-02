package main

func Example_selectionSortA() {
	selectionSort([]int{4, 2, 3, 1})
	// Output:
	// [4 2 3 1]
	// [1 2 3 4]
}
func Example_selectionSortB() {
	selectionSort([]int{2, 3, 4, 1})
	// Output:
	// [2 3 4 1]
	// [1 3 4 2]
	// [1 2 4 3]
	// [1 2 3 4]
}
func Example_selectionSortC() {
	selectionSort([]int{1, 2, 4, 3})
	// Output:
	// [1 2 4 3]
	// [1 2 3 4]
}

func Example_insertionSortA() {
	insertionSort([]int{4, 2, 3, 1})
	// Output:
	// [4 2 3 1]
	// [2 4 3 1]
	// [2 3 4 1]
	// [1 2 3 4]
}
func Example_insertionSortB() {
	insertionSort([]int{2, 3, 4, 1})
	// Output:
	// [2 3 4 1]
	// [1 2 3 4]
}
func Example_insertionSortC() {
	insertionSort([]int{1, 2, 4, 3})
	// Output:
	// [1 2 4 3]
	// [1 2 3 4]
}

func Example_bubbleSortA() {
	bubbleSort([]int{4, 2, 3, 1})
	// Output:
	// [4 2 3 1]
	// [4 2 1 3]
	// [4 1 2 3]
	// [1 4 2 3]
	// [1 2 4 3]
	// [1 2 3 4]
}
func Example_bubbleSortB() {
	bubbleSort([]int{2, 3, 4, 1})
	// Output:
	// [2 3 4 1]
	// [2 3 1 4]
	// [2 1 3 4]
	// [1 2 3 4]
}
func Example_bubbleSortC() {
	bubbleSort([]int{1, 2, 4, 3})
	// Output:
	// [1 2 4 3]
	// [1 2 3 4]
}
