package main

import "fmt"

// ※ k桁の数値を対象(k個の数値の組とみなしてソート)．
// 1. 10個の容器(バケツ)を用意(a[0]〜a[9])．
// 2. 1番目からk番目まで(最下位桁である1の位から最上位桁まで)3.を繰り返す．
// 3. 当該要素をキーとして，用意した容器でバケットソート法を行う．

const digit = 3 //3桁までの自然数が対象

var nums = []int{374, 889, 309, 397, 987, 473, 346, 607, 881} //配列データ

// 基数ソート法：radix sort
func main() {
	fmt.Println(nums) //ソート前の配列を表示

	for i := 1; i <= digit; i++ {
		a := [10][]int{} // 10個の容器を用意(a[0]〜a[9])

		//各数値を対応する容器に格納
		for _, num := range nums {
			x := getSpecifiedDigit(num, i) //numのi桁目の値を取得
			a[x] = append(a[x], num)       //numをa[x]に格納
		}

		//容器の中身を順番に取り出す
		count := 0
		for _, bucket := range a {
			for _, v := range bucket {
				nums[count] = v
				count++
			}
		}
		fmt.Println(nums) //i桁目までソート後の配列を表示
	}
}

func getSpecifiedDigit(num, i int) int {
	switch i {
	case 1:
		return num % 10
	case 2:
		return num % 100 / 10
	case 3:
		return num % 1000 / 100
	}
	return -1 //エラー
}
