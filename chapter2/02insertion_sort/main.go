package main

import "fmt"

// 1. 変数iを挿入する要素の配列番号として扱い，2.〜4.の操作を，i=1〜n-1
//    について繰り返す．
// 2. すでにソート済みのa[0]〜a[i-1]の中で，a[i]が入るべき位置kを見つける
//    (a[i]が最大のときはk=i)．
// 3. a[k]〜a[i-1]を一つ後ろにずらす(k=iのときは何もしない)．
// 4. a[i]の値を，空いたa[k]に代入する．

const n = 8 //配列の要素数

var a = [n]int{7, 2, 3, 4, 8, 1, 5, 6} //配列データ

// 単純挿入法：insertion sort
func main() {
	fmt.Println(a) //ソート前の配列を表示

	for i := 1; i < n; i++ {
		temp := a[i]
		j := i
		//GoにWhile文は存在しないので，for文で代用
		for j > 0 && a[j-1] > temp { //挿入すべき位置を見つける
			a[j] = a[j-1] //一つ後ろにずらす
			j = j - 1
		}
		a[j] = temp //ここに挿入する

		fmt.Println(a) //挿入後の配列を表示
	}
}
