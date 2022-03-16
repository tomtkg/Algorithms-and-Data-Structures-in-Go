package main

// 1. iを先頭要素の番号0，jを末尾要素の番号n-1とする．
// 2. 配列から任意の要素を取り出し，これを基準値pivotとする．
// 3. i > jとなるまで，4.〜6.を繰り返す．
// 4. iを増加させていき，要素a[i] >= pivotが見つかるまで配列を走査する．
// 5. jを減少させていき，要素a[i] <= pivotが見つかるまで配列を走査する．
// 6. i <= jならば，a[i]とa[j]を交換し，iを1増加，jを1減少させて3.に戻る．

//クイックソート法：quick sort
func quickSortA(seq []int, start, end int) {
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
		quickSortA(seq, start, j)
	}
	if end > i { //pivotより後をクイックソート
		quickSortA(seq, i, end)
	}
}
