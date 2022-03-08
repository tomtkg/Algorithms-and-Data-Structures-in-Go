package main

import "fmt"

// 1. 変数iを最小要素を入れる配列番号として扱い，2.の操作を，i=0〜n-2
//    について繰り返す．
// 2. a[i]〜a[n-1]の中から，最小要素a[k]を探し，a[i]と入れ替える．

const n = 10 //配列の要素数

var a = [n]int{6, 9, 12, 7, 15, 23, 2, 10, 4, 20} //配列データ
//var a = [n]int{13, 22, 12, 10, 29, 14, 20, 18, 24, 28} //Q2.1

//単純選択法：selection sort
func main() {
	fmt.Println(a) //ソート前の配列を表示

	//単純選択法でソート
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if a[k] > a[j] { //最少要素を検索
				k = j
			}
		}
		a[i], a[k] = a[k], a[i] //検索範囲の最初の要素と交換

		fmt.Println(a) //最少要素選択後の配列を表示
	}
}
