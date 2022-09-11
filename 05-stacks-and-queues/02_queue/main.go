package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const QUEUESIZE = 10 //キューサイズ

var queue [QUEUESIZE]float64 //キュー用配列
var head int                 //ヘッドの初期化
var tail int                 //テールの初期化

func initQueue() {
	queue = [QUEUESIZE]float64{} //キュー用配列
	head = 0                     //ヘッドの初期化
	tail = 0                     //テールの初期化
}

// データをキューにenqueueする関数
func enqueue(data float64) int {
	if tail >= QUEUESIZE { //データの末尾が，キューの最後に達しているかどうか
		return 0
	} else {
		queue[tail] = data //テールの示す位置にデータを格納し，
		tail++             //テールを1個分進める
		return 1
	}
}

// キューからデータをdequeueする関数
func dequeue(data *float64) int {
	if head >= tail { //キューにデータが残っているかどうか
		return 0
	} else {
		*data = queue[head] //ヘッドの示す位置からデータを取得し，
		head++              //ヘッドを1個分進める
		return 1
	}
}

func main() {
	seq := [][]float64{} //データ格納用配列
	_ = json.NewDecoder(os.Stdin).Decode(&seq)

	initQueue()
	fmt.Println(queue, head, tail)
	for _, nums := range seq {
		for _, v := range nums {
			if enqueue(v) == 0 {
				panic("Enqueue error")
			}
		}
		fmt.Println(queue, head, tail)
		for {
			var data float64
			if dequeue(&data) == 0 {
				break
			}
			fmt.Println(data)
		}
		fmt.Println(queue, head, tail)
	}
}
