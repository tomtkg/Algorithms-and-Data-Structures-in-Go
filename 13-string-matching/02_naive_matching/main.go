package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	problems := [][2]string{}
	_ = json.NewDecoder(os.Stdin).Decode(&problems)

	for _, p := range problems {
		fmt.Printf("text = %s\n", p[0])
		fmt.Printf("word = %s\n", p[1])
		ans := naiveMatching([]rune(p[0]), []rune(p[1]))
		fmt.Println(ans)
	}
}

func naiveMatching(s, w []rune) int {
	for i, j := 0, 0; i+j < len(s); {
		if s[i+j] == w[j] {
			j++
			if j == len(w) {
				return i
			}
		} else { //単語を1文字ずらして，最初から照合しなおす
			i++
			j = 0
		}
	}
	return len(s) //最後まで探索して，照合に失敗した
}
