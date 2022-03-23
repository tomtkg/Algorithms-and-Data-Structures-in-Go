package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	ans := calc(s)
	fmt.Printf("%.2f\n", ans)
}

func calc(s *bufio.Scanner) float64 {
	for s.Scan() {
		str := s.Text()
		//読み込んだ文字列が，演算子を示す文字列か？
		if data, err := strconv.ParseFloat(str, 64); err != nil {
			//演算子の場合
			data1 := calc(s)
			data2 := calc(s)
			switch str {
			case "+": //加算
				return (data1 + data2)
			case "-": //減算
				return (data1 - data2)
			case "*": //乗算
				return (data1 * data2)
			case "/": //除算
				return (data1 / data2)
			default:
				panic("Unknown operator")
			}
		} else {
			//数字の場合
			return data
		}
	}
	panic("not enough operand")
}
