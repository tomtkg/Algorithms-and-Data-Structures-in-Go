package main

import "fmt"

// ※ 0〜n-1までの各整数値を対象．
// 1. n個の容器(バケツ)を用意(a[0]〜a[n-1])
// 2. 各数値を対応する容器に順番に格納(a[x]=x，同一の場合，元の順番を維持する)．
// 3. 容器に入っている数値を順番に出力．

const N = 4 //配列の要素数
const M = 5 //容器の個数

//バケットソート法：bucket sort
func main() {
	a := [N]int{3, 1, 4, 0}
	b := [M]int{}

	//容器の初期化
	for i := range b {
		b[i] = -1
	}

	//ソート前の配列を表示
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()

	//各数値を対応する容器に格納
	for _, v := range a {
		b[v] = v
	}

	//ソート後の配列を表示
	for _, v := range b {
		if v > -1 {
			fmt.Print(v) //データが格納されている容器の値のみを出力
		}
	}
	fmt.Println()
}
