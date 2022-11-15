package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 下降修復
func downheap(data []int, v, n int) {
	if v > n/2 { //vが子ノードをもたなければ何もしない
		return
	}
	w := 2 * v //左の子ノードをwとする
	if w < n && data[w] < data[w+1] {
		w++ //右の子ノードをwとする
	}
	if data[v] < data[w] {
		data[v], data[w] = data[w], data[v]
		downheap(data, w, n)
	}
}

func main() {
	var data []int
	_ = json.NewDecoder(os.Stdin).Decode(&data)

	fmt.Println("** Before heap...")
	fmt.Println(data)
	n := len(data) - 1

	for i := n / 2; i > 0; i-- {
		downheap(data, i, n)
	}

	fmt.Println("** After heap...")
	fmt.Println(data)

	for i := n; i > 0; i-- {
		data[1], data[i] = data[i], data[1]
		downheap(data, 1, i-1)
	}

	fmt.Println("** After sort...")
	fmt.Println(data)
}
