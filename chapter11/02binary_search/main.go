package main

import (
	"fmt"
)

const n = 8                                   //配列の要素数
var a = []int{10, 14, 19, 23, 27, 31, 34, 38} //配列データ

func main() {
	fmt.Println(binSearch(23))
	fmt.Println(binSearch(16))
}

func binSearch(x int) int {
	// 前処理
	if x <= a[0] {
		return 0
	} else if a[n-1] < x {
		return n
	}

	left, right := 0, n-1
	for right-left > 1 {
		middle := (right + left) / 2
		if x <= a[middle] {
			right = middle
		} else {
			left = middle
		}
	}
	return right
}
