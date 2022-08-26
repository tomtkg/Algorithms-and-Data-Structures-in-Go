package main

import "fmt"

func main() {
	result := sumA(10)
	fmt.Println("sumA:", result)
	result = sumB(10)
	fmt.Println("sumB:", result)
}

// 整数1〜nの総和を求める．
func sumA(n int) int {
	s := 0                    //データ
	for i := 1; i <= n; i++ { //アルゴリズム
		s += i
	}
	return s
}

// 1〜nの和を二つ用意する．ただし，一つは逆順に並べ，それぞれの項を加算する．
// それぞれの加算結果はn+1となり，これがn個存在することになる．
// しかし，これは最初に1〜nの和を二つ用意した結果なので，2で割る．
func sumB(n int) int {
	s := 0              //データ
	s = n * (n + 1) / 2 //アルゴリズム
	return s
}
