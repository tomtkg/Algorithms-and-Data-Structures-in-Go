package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 1. iを先頭要素の番号0，jを末尾要素の番号n-1とする．
// 2. 配列から任意の要素を取り出し，これを基準値pivotとする．
// 3. i > jとなるまで，4.〜6.を繰り返す．
// 4. iを増加させていき，要素a[i] >= pivotが見つかるまで配列を走査する．
// 5. jを減少させていき，要素a[i] <= pivotが見つかるまで配列を走査する．
// 6. i <= jならば，a[i]とa[j]を交換し，iを1増加，jを1減少させて3.に戻る．

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
	i, j := start, end    //初期走査位置の設定
	pivot := seq[(i+j)/2] //pivotの設定
	for i <= j {
		for ; seq[i] < pivot; i++ { //pivotより大きい値を走査
		}
		for ; seq[j] > pivot; j-- { //pivotより小さい値を走査
		}
		if i <= j {
			seq[i], seq[j] = seq[j], seq[i] //値の交換
			i, j = i+1, j-1
		}
	}
	if start < j { //pivotより前をクイックソート
		quick(seq, start, j)
	}
	if end > i { //pivotより後をクイックソート
		quick(seq, i, end)
	}
}
