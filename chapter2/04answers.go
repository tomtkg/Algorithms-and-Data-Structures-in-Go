package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func main() {
	a := []int{6, 9, 3, 8, 7, 5, 4, 2, 1} // 配列データ

	for _, f := range []func(a []int){
		selectionSort,
		insertionSort,
		bubbleSort,
	} {
		printFunctionName(f)
		copy := append([]int{}, a...) //deep copy
		fmt.Println(copy)             //ソート前の配列を表示
		f(copy)
	}
}

func printFunctionName(f interface{}) {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	arr := strings.Split(name, ".")
	fmt.Println(arr[len(arr)-1])
}

func selectionSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[k] > a[j] { //最少要素を検索
				k = j
			}
		}
		if k != i {
			a[i], a[k] = a[k], a[i] //検索範囲の最初の要素と交換
			fmt.Println(a)          //最少要素選択後の配列を表示
		}
	}
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i
		//GoにWhile文は存在しないので，for文で代用
		for j > 0 && a[j-1] > temp { //挿入すべき位置を見つける
			a[j] = a[j-1] //一つ後ろにずらす
			j = j - 1
		}
		if j != i {
			a[j] = temp    //ここに挿入する
			fmt.Println(a) //挿入後の配列を表示
		}
	}
}

func bubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := len(a) - 1; j > i; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				fmt.Println(a) //交換後の配列を表示
			}
		}
	}
}
