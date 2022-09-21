package main

import "fmt"

// 1. 探索する値が配列の数値の範囲外でないか確認する (実装例の前処理を参照)．
// 2. left=0，right=1とする (探索範囲をa[0]〜a[n-1]とする)．
// 3. right-left>1となる間，4. 5. を繰り返す．
// 4. middle=(left+right)/2 (切り捨て)とする．
// 5. x≦a[middle]ならば，right=middleとし，そうでなければleft=middleとする．
// 6. rightの値を出力する．

const n = 8                                   //配列の要素数
var a = []int{10, 14, 19, 23, 27, 31, 34, 38} //配列データ

func main() {
	fmt.Println("x=23")
	binSearch(23)
	fmt.Println("x=16")
	binSearch(16)
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
		fmt.Printf("l: a[%d]=%d, m: a[%d]=%d, r: a[%d]=%d\n",
			left, a[left], middle, a[middle], right, a[right])
		if x <= a[middle] {
			right = middle
		} else {
			left = middle
		}
	}
	fmt.Printf("l: a[%d]=%d, r: a[%d]=%d\n",
		left, a[left], right, a[right])
	return right
}
