package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1013)
	nums := [][2]int{{372, 377}, {380, 382}, {402, 410}, {712, 715}, {740, 746}, {813, 816}}
	for _, v := range nums {
		for i := v[0]; i <= v[1]; i++ {
			a[i] = -1
		}
	}
	f := []func(i int, s string) int{
		func(i int, s string) int { return (i + 1) % 1013 },
		func(i int, s string) int { return (i + 5*int(s[1]) + 2*len(s)) % 1013 },
	}
	a[409], a[743] = 1, 1
	for j, s := range []string{"MOTHER", "PARENT"} {
		i := 2*int(s[0]) + 3*int(s[2])
		for ; a[i] != 1; i = f[j](i, s) {
			fmt.Printf("%d->", i)
		}
		fmt.Println(i)
	}
	for _, ff := range f {
		b := append([]int{}, a...)
		for _, s := range []string{"STEAK", "HAMBURG", "HUNGRY"} {
			i := 2*int(s[0]) + 3*int(s[2])
			for ; b[i] != 0; i = ff(i, s) {
			}
			b[i] = 1
			fmt.Printf(" %s:%d,", s, i)
		}
		fmt.Println()
	}
}
