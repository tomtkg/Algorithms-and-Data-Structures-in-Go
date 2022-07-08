package main

import "fmt"

func main() {
	s := []rune("algorithm")
	w := []rune("go")
	fmt.Printf("text = %c\n", s)
	fmt.Printf("word = %c\n", w)

	ans := naiveMatching(s, w)
	fmt.Println(ans)
}

func naiveMatching(s, w []rune) int {
	for i, j := 0, 0; i+j < len(s); {
		if s[i+j] == w[j] {
			j++
			if j == len(w) {
				return i
			}
		} else {
			i++
			j = 0
		}
	}
	return len(s)
}
