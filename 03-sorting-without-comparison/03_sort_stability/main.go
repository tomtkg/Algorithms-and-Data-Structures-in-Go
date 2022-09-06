package main

import "fmt"

type table struct {
	key  int
	data string
}

func main() {
	t := []table{
		{120, "みかん"},
		{110, "りんご"},
		{90, "キウイ"},
		{120, "桃"},
		{50, "いちご"},
	}
	fmt.Println("ソート前の配列:", t)
	copy := append([]table{}, t...)
	selectionSort(copy)
	fmt.Println("単純選択法の結果:", copy)
	copy = append([]table{}, t...)
	bubbleSort(copy)
	fmt.Println("バブルソート法の結果:", copy)
}

func selectionSort(t []table) {
	for i := 0; i < len(t)-1; i++ {
		k := i
		for j := i + 1; j < len(t); j++ {
			if t[k].key > t[j].key { //最少要素を検索
				k = j
			}
		}
		if k != i {
			t[i], t[k] = t[k], t[i] //検索範囲の最初の要素と交換
		}
	}
}

func bubbleSort(t []table) {
	for i := 0; i < len(t)-1; i++ {
		for j := len(t) - 1; j > i; j-- {
			if t[j-1].key > t[j].key {
				t[j-1], t[j] = t[j], t[j-1]
			}
		}
	}
}
