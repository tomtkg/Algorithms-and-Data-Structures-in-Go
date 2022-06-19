package main

import (
	"fmt"
)

func main() {
	hashing(5, 10, []int{11, 3, 32, 4, 50})
	hashing(5, 10, []int{11, 3, 32, 4, 50, 34})
	hashing(9, 9, []int{35, 21, 103, 96, 232, 36, 185})
	hashing(9, 9, []int{29, 165, 333, 8, 97, 56, 11})
}
func hashing(n, v int, nums []int) {
	a := make([]int, n)
	for _, num := range nums {
		a[num%v] = num
	}
	fmt.Println(a)
}
