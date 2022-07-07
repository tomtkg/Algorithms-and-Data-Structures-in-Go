package main

import "fmt"

func main() {
	s := []rune("algorithm")
	for _, x := range []rune{'t', 'r'} {
		fmt.Printf("x = %c\n", x)
		ans := charMatching(s, x)
		fmt.Println(ans)
	}
}

func charMatching(s []rune, x rune) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return len(s)
}
