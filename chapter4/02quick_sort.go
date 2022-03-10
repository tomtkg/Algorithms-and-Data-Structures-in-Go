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
	fmt.Println(seq)          //ソート前の配列を表示
	quick(seq, 0, len(seq)-1) //クイックソート
	fmt.Println(seq)          //ソート後の配列を表示
}

//クイックソート法：quick sort
func quick(seq []int, start, end int) {
	if start < end {
		i := start                 //pivotの位置（仮）
		pivot := seq[end]          //配列の末尾要素をpivotに設定
		for j := i; j < end; j++ { //pivotより小さい値を走査
			if seq[j] < pivot {
				seq[i], seq[j] = seq[j], seq[i] //値の交換
				i++                             //pivotの位置（仮）を変更
			}
		}
		seq[i], seq[end] = seq[end], seq[i] //pivotを適切な位置に置く
		quick(seq, start, i-1)              //pivotより前をクイックソート
		quick(seq, i+1, end)                //pivotより後をクイックソート
	}
}
