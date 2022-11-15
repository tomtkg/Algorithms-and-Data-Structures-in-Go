package main

import "fmt"

// 上昇修復
// ノードvの値がその親uの値より大きければ，uとvの値を交換し，
// 次にノードuに対して上昇修復を繰り返す．
func upheap(data []int, v int) {
	if v == 1 {
		return
	}
	w := v / 2
	if data[v] > data[w] {
		data[v], data[w] = data[w], data[v]
		upheap(data, w)
	}
}

func main() {
	data := []int{0, 15, 7, 9, 4}

	fmt.Println("** Before heap...")
	fmt.Println(data)

	data = append(data, 24)
	upheap(data, len(data)-1)

	fmt.Println("** After heap...")
	fmt.Println(data)
}
