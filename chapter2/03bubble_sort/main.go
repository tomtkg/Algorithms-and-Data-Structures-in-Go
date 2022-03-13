package main

import "fmt"

// 1. i=0〜n-2について，2.〜4.を繰り返す．
// 2. jをn-1とする．
// 3. a[j-1]とa[j]を比較し，a[j]のほうが小さければ，a[j-1]とa[j]
//    を交換する．
// 4. jを1減らす．減らしたjに対してj>iであれば，3.に戻る．

const n = 10 //配列の要素数

var a = [n]int{6, 9, 12, 7, 15, 23, 2, 10, 4, 20} //配列データ

// バブルソート法：bubble sort
func main() {
	fmt.Println(a) //ソート前の配列を表示

	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]

				fmt.Println(a) //交換後の配列を表示
			}
		}
	}
}
