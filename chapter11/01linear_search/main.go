package main

import (
	"fmt"
)

const n = 8                                   //配列の要素数
var a = []int{28, 14, 15, 29, 27, 23, 13, 30} //配列データ

func main() {
	var i int //制御変数
	x := 23   //探索するデータ
	for i = 0; i < n && x != a[i]; i++ {
	}
	fmt.Println(i)
}
