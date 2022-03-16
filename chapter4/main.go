package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	seq := []int{}     //データ格納用配列
	reader := os.Stdin //読み込み用のファイルポインタ
	if len(os.Args) > 1 {
		if r, err := os.Open(os.Args[1]); err == nil {
			reader = r
		}
	}
	//ファイルからのデータの読み込み
	if err := json.NewDecoder(reader).Decode(&seq); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, f := range []func([]int, int, int){quickSortA, quickSortB} {
		copy := append([]int{}, seq...)
		fmt.Println(copy)      //ソート前の配列を表示
		f(copy, 0, len(seq)-1) //クイックソート
		fmt.Println(copy)      //ソート後の配列を表示
	}
}
