package main

import "fmt"

var table [20]int

func main() {
	s := []rune("algorithm")
	w := []rune("go")
	fmt.Printf("text = %c\n", s)
	fmt.Printf("word = %c\n", w)

	makePatternMatchingTable(w)
	ans := kmpMatching(s, w)
	fmt.Println(ans)
}

func kmpMatching(s, w []rune) int {
	for i, j := 0, 0; i+j < len(s); {
		if s[i+j] == w[j] {
			j++
			if j == len(w) {
				return i
			}
		} else {
			i += j - table[j]
			if j > 0 {
				j = table[j]
			}
		}
	}
	return len(s)
}

func makePatternMatchingTable(w []rune) {
	for j := range w {
		k := j - 1
		for k > 0 && !equal(w, j, k) {
			k--
		}
		table[j] = k
	}
}

func equal(w []rune, j, k int) bool {
	for i := 0; i < k; i++ {
		if w[i] != w[j-k+i] {
			return false
		}
	}
	return true
}
