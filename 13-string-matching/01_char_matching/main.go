package main

import "fmt"

func main() {
	s := []rune("algorithm")
	for _, x := range []rune{'t', 'r', 'a', 'z'} {
		fmt.Printf("x = %c\n", x)
		ans := charMatching(s, x)
		fmt.Println(ans)
	}
}

func charMatching(s []rune, x rune) int {
	for i := range s { //1. I = 0〜N-1まで，2. を繰り返す．
		if s[i] == x { //2. s[I] = xならば，iを出力して終了する．
			return i
		}
	}
	return len(s)
}
