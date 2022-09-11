package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var m = map[string]func(float64, float64) float64{
	"+": func(x, y float64) float64 { return x + y }, //加算
	"-": func(x, y float64) float64 { return x - y }, //減算
	"*": func(x, y float64) float64 { return x * y }, //乗算
	"/": func(x, y float64) float64 { return x / y }, //除算
}

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
			if f, ok := m[str]; ok {
				return f(data1, data2)
			} else {
				panic("Unknown operator")
			}
		} else {
			//数字の場合
			return data
		}
	}
	panic("not enough operand")
}
