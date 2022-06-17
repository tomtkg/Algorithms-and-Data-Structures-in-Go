package main

import (
	"fmt"
)

func main() {
	a := []int{15, 20, 33, 41, 59, 68, 72, 75, 90}
	for _, v := range []int{75, 20, 69, 100} {
		fmt.Printf(",戻り値:%d\n", binSearch(a, len(a), v))
	}
}

func binSearch(a []int, n, x int) int {
	l, r := 0, n-1
	fmt.Printf("rの値の変化:%d", r)

	if x <= a[0] {
		return 0
	} else if a[n-1] < x {
		return n
	}

	for r-l > 1 {
		middle := (r + l) / 2
		if x <= a[middle] {
			r = middle
			fmt.Printf("->%d", r)
		} else {
			l = middle
		}
	}
	return r
}
