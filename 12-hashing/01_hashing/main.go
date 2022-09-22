package main

import "fmt"

func hashing(n, v int, record []int) {
	H := func(x int) int { return x % v } // ハッシュ関数H
	a := make([]int, n)                   // 1. データを格納する配列の要素数nを設定する．
	for _, k := range record {
		i := H(k) //2. キーkに対して，ハッシュ関数Hを適用し，値iを算出する (H(k)=i)．
		a[i] = k  //3. キーkをa[i]に格納する
	}
	fmt.Println(a)
}

func main() {
	fmt.Println("12.2節")
	hashing(5, 10, []int{11, 3, 32, 4, 50})
	hashing(5, 10, []int{11, 3, 32, 4, 50, 34})
	fmt.Println("問12.1(1)")
	hashing(9, 9, []int{35, 21, 103, 96, 232, 36, 185})
	fmt.Println("問12.1(2)")
	hashing(9, 9, []int{29, 165, 333, 8, 97, 56, 11})
}
