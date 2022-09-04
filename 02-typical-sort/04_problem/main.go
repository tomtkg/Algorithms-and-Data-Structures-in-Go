package main

import (
	"fmt"
	"mylib"
)

var num int

func print(x interface{}) {
	fmt.Printf("(%d) %v\n", num, x)
	num++
}

func main() {
	functions := []func(a []int){
		selectionSort,
		insertionSort,
		bubbleSort,
	}
	problems := []mylib.Tuple[string, []int]{
		{V1: "問2.", V2: []int{13, 22, 12, 10, 29, 14, 20, 18, 24, 28}},
		{V1: "章末問題", V2: []int{6, 9, 3, 8, 7, 5, 4, 2, 1}},
	}
	for _, p := range problems {
		for i, f := range functions {
			num = 0
			fmt.Printf("%s%d\n", p.V1, i+1)
			f(append([]int{}, p.V2...))
		}
	}
}

func selectionSort(a []int) {
	print(a) //ソート前の配列を表示
	for i := 0; i < len(a)-1; i++ {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[k] > a[j] { //最少要素を検索
				k = j
			}
		}
		if k != i {
			a[i], a[k] = a[k], a[i] //検索範囲の最初の要素と交換
			print(a)                //最少要素選択後の配列を表示
		}
	}
}

func insertionSort(a []int) {
	print(a) //ソート前の配列を表示
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i
		//GoにWhile文は存在しないので，for文で代用
		for j > 0 && a[j-1] > temp { //挿入すべき位置を見つける
			a[j] = a[j-1] //一つ後ろにずらす
			j = j - 1
		}
		if j != i {
			a[j] = temp //ここに挿入する
			print(a)    //挿入後の配列を表示
		}
	}
}

func bubbleSort(a []int) {
	print(a) //ソート前の配列を表示
	for i := 0; i < len(a)-1; i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				print(a) //交換後の配列を表示
			}
		}
	}
}
