package main

import (
	"fmt"
	"mylib"
)

// 1. iを先頭要素の番号0，jを末尾要素の番号n-1とする．
// 2. 配列から任意の要素を取り出し，これを基準値pivotとする．
// 3. i > jとなるまで，4.〜6.を繰り返す．
// 4. iを増加させていき，要素a[i] >= pivotが見つかるまで配列を走査する．
// 5. jを減少させていき，要素a[i] <= pivotが見つかるまで配列を走査する．
// 6. i <= jならば，a[i]とa[j]を交換し，iを1増加，jを1減少させて3.に戻る．

// クイックソート法：quicksort
func quicksortA(seq []int, start, end int) {
	i, j := start, end    //初期走査位置の設定
	pivot := seq[(i+j)/2] //pivotの設定
	for i <= j {
		for ; seq[i] < pivot; i++ { //pivotより大きい値を走査
		}
		for ; seq[j] > pivot; j-- { //pivotより小さい値を走査
		}
		if i <= j {
			seq[i], seq[j] = seq[j], seq[i] //値の交換
			i++
			j--

		}
	}
	if start < j { //pivotより前をクイックソート
		quicksortA(seq, start, j)
	}
	if end > i { //pivotより後をクイックソート
		quicksortA(seq, i, end)
	}
}

// クイックソート法：quicksort
func quicksortB(seq []int, start, end int) {
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
		quicksortB(seq, start, i-1)         //pivotより前をクイックソート
		quicksortB(seq, i+1, end)           //pivotより後をクイックソート
	}
}

func main() {
	funcs := []func([]int, int, int){quicksortA, quicksortB}
	for _, seq := range [][]int{
		{37, 33, 29, 25, 22, 20, 19, 14, 10, 8},
		{6, 9, 12, 7, 15, 23, 2, 10, 4, 20},
		{6, 9, 3, 8, 7, 5, 4, 2, 1},
	} {
		for _, f := range funcs {
			mylib.PrintFuncName(f)
			copy := append([]int{}, seq...)
			fmt.Println(copy)      //ソート前の配列を表示
			f(copy, 0, len(seq)-1) //クイックソート
			fmt.Println(copy)      //ソート後の配列を表示
		}
	}
}
