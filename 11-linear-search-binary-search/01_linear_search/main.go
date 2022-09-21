package main

import "fmt"

// 1. i=0〜n-1に対して，次の2.を繰り返す．
// 2. a[i]=xならば，iを出力して終了する．
// 3. nを出力する．

const n = 8                                   //配列の要素数
var a = []int{28, 14, 15, 29, 27, 23, 13, 30} //配列データ

func main() {
	var i int //制御変数
	x := 23   //探索するデータ
	for i = 0; i < n && x != a[i]; i++ {
	}
	fmt.Println(i)
}
