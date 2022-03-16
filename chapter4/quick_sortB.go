package main

//クイックソート法：quick sort
func quickSortB(seq []int, start, end int) {
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
		quickSortB(seq, start, i-1)         //pivotより前をクイックソート
		quickSortB(seq, i+1, end)           //pivotより後をクイックソート
	}
}
