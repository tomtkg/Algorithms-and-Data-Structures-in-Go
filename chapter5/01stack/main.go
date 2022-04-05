package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const STACKSIZE = 10 //スタックサイズ

var stack [STACKSIZE]float64 //スタック用配列
var sp int                   //スタックポインタの初期化

func initStack() {
	stack = [STACKSIZE]float64{} //スタック用配列
	sp = 0                       //スタックポインタの初期化
}

//データをスタックにpushする関数
func push(data float64) int {
	if sp >= STACKSIZE { //スタックポインタの示す位置が範囲内かどうか
		return 0
	} else {
		stack[sp] = data //スタックポインタの示す位置にデータを格納し，
		sp++             //スタックポインタを1個分進める
		return 1
	}
}

//スタックからデータをpopする関数
func pop(data *float64) int {
	if sp <= 0 { //スタックが空かどうか
		return 0
	} else {
		sp--              //スタックポインタを1個分戻し，
		*data = stack[sp] //スタックポインタの示す位置からデータを取得する
		return 1
	}
}

func main() {
	var data float64         //入力データ
	var data1, data2 float64 //計算用
	initStack()
	s := bufio.NewScanner(os.Stdin)

	//EOFまで入力を繰り返し
	for s.Scan() {
		str := s.Text()
		//読み込んだ文字列が，演算子を示す文字列か？
		if data, err := strconv.ParseFloat(str, 64); err != nil {
			//演算子の場合，スタックのデータ二つを取り出す
			if pop(&data2) == 1 && pop(&data1) == 1 {
				switch str {
				case "+": //加算
					push(data1 + data2)
				case "-": //減算
					push(data1 - data2)
				case "*": //乗算
					push(data1 * data2)
				case "/": //除算
					push(data1 / data2)
				default:
					fmt.Fprintln(os.Stderr, "Unknown operator")
					os.Exit(1)
				}
			} else {
				fmt.Fprintln(os.Stderr, "Stack underflow")
				os.Exit(1)
			}
		} else {
			//数字の場合，データをスタックに入れる
			if push(data) == 0 {
				fmt.Fprintln(os.Stderr, "Stack overflow")
				os.Exit(1)
			}
		}
	}

	//スタックに残った最後の数字が演算結果
	if pop(&data) == 1 {
		fmt.Printf("%.2f\n", data)
	} else {
		fmt.Fprintln(os.Stderr, "Stack underflow")
	}
}
