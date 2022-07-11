package main

import "fmt"

var table [20]int

func main() {
	s := []rune("cucutexecutexecution")
	w := []rune("cutexecution")

	makePatternMatchingTable(w)
	for j, v := range w {
		fmt.Printf("%d %c %d %d\n", j, v, table[j], j-table[j])
	}
	fmt.Println()
	kmpMatching(s, w)
}

func kmpMatching(s, w []rune) int {
	for i, j := 0, 0; i+j < len(s); {
		fmt.Println(i, j)
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
