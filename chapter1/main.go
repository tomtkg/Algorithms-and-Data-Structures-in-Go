package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "a":
			a()
		case "b":
			b()
		}
	}
}

// 整数1〜nの総和を求める．
func a() {
	n, s := 10, 0             //データ
	for i := 1; i <= n; i++ { //アルゴリズム
		s += i
	}
	fmt.Println(s)
}

// 1〜nの和を二つ用意する．ただし，一つは逆順に並べ，それぞれの項を加算する．
// それぞれの加算結果はn+1となり，これがn個存在することになる．
// しかし，これは最初に1〜nの和を二つ用意した結果なので，2で割る．
func b() {
	n, s := 10, 0       //データ
	s = n * (n + 1) / 2 //アルゴリズム
	fmt.Println(s)
}
