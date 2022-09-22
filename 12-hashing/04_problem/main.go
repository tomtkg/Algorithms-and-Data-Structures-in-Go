package main

import "fmt"

func main() {
	//定義
	n := 1013
	H := func(s string) int { return 2*int(s[0]) + 3*int(s[2]) }
	DH := func(s string) int { return 5*int(s[1]) + 2*len(s) }
	a := make([]int, n)

	//初期状態
	nums := [][2]int{{372, 377}, {380, 382}, {402, 410}, {712, 715}, {740, 746}, {813, 816}}
	for _, v := range nums {
		for i := v[0]; i <= v[1]; i++ {
			a[i] = -1
		}
	}

	f := []func(i int, s string) int{
		func(i int, _ string) int { return (i + 1) % n },     //線形探査法
		func(i int, s string) int { return (i + DH(s)) % n }, //ダブルハッシュ法
	}

	fmt.Println("(1)")
	a[409], a[743] = 1, 1
	for j, s := range []string{"MOTHER", "PARENT"} {
		i := H(s)
		for ; a[i] != 1; i = f[j](i, s) {
			fmt.Printf("%d -> ", i)
		}
		fmt.Println(i)
	}

	fmt.Println("(2)")
	for j := range f {
		b := append([]int{}, a...)
		for _, s := range []string{"STEAK", "HAMBURG", "HUNGRY"} {
			i := H(s)
			for ; b[i] != 0; i = f[j](i, s) {
			}
			b[i] = 1
			fmt.Printf(" %s: %d,", s, i)
		}
		fmt.Println()
	}
}
