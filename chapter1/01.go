package main

import "fmt"

// 整数1〜nの総和を求める．

func main() {
	n, s := 10, 0             //データ
	for i := 1; i <= n; i++ { //アルゴリズム
		s = s + i
	}
	fmt.Println(s)
}
