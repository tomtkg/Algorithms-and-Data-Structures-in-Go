package main

import "fmt"

func hashing(n, v int, record []int) {
	H := func(x int) int { return x % v }     // ハッシュ関数H
	DH := func(x int) int { return x/10 + 1 } // ハッシュ関数DH
	a := make([]int, n)
	for _, k := range record {
		i := H(k)
		for ; a[i] != 0; i = (i + DH(k)) % n {
		}
		a[i] = k
	}
	fmt.Println(a)
}

func main() {
	fmt.Println("12.3節")
	hashing(11, 10, []int{11, 3, 32, 4, 50, 34, 19, 39})
	fmt.Println("問12.2(1)")
	hashing(11, 10, []int{11, 3, 32, 4, 50, 16, 36})
	fmt.Println("問12.2(2)")
	// hashing(8, 10, []int{11, 3, 32, 4, 50, 16, 36}) error
}
