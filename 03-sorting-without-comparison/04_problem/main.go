package main

import "fmt"

const digit = 3 //3桁までの自然数が対象

// 基数ソート法：radix sort
func main() {
	nums := []int{523, 39, 785, 257, 351, 616, 292, 788, 802}
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
	if num < 0 {
		num = -num
	}
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
