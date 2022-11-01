package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var table [20]int

func main() {
	problems := [][2]string{}
	_ = json.NewDecoder(os.Stdin).Decode(&problems)

	for _, p := range problems {
		fmt.Printf("text = %s\n", p[0])
		fmt.Printf("word = %s\n", p[1])
		makePatternMatchingTable([]rune(p[1]))
		ans := kmpMatching([]rune(p[0]), []rune(p[1]))
		fmt.Println(ans)
	}
}

func kmpMatching(s, w []rune) int {
	for i, j := 0, 0; i+j < len(s); {
		if s[i+j] == w[j] {
			j++
			if j == len(w) {
				return i
			}
		} else { //単語を指定された分だけずらして，照合しなおす
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
