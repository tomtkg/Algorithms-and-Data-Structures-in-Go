package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const NODES = 9 //ノード数
const START = 0 //探索開始ノード

type edge struct {
	startNode int //辺の一端のノード番号
	endNode   int //辺の一端のノード番号
}

var (
	matrix  [NODES][NODES]int //隣接行列
	bfFlag  [NODES]int        //探索状況（0: 未探索，1: 探索済）
	bfTree  [NODES - 1]edge   //探索木
	edgeCnt int
)

func bfSearch(start int) {
	queue := [NODES]int{start}
	bfFlag[start] = 1

	for head, tail := 0, 1; head < tail; head++ {
		u := queue[head]
		for v := 0; v < NODES; v++ {
			if matrix[u][v] == 1 && bfFlag[v] == 0 {
				bfTree[edgeCnt].startNode = u
				bfTree[edgeCnt].endNode = v
				edgeCnt++
				queue[tail] = v
				tail++
				bfFlag[v] = 1
			}
		}
	}
}
func main() {
	graph := [][2]int{}
	_ = json.NewDecoder(os.Stdin).Decode(&graph)

	//初期化
	matrix = [NODES][NODES]int{}
	bfFlag = [NODES]int{}
	edgeCnt = 0

	for _, v := range graph {
		matrix[v[0]][v[1]] = 1
		matrix[v[1]][v[0]] = 1
	}

	for _, v := range matrix { //隣接行列を表示する
		fmt.Println(v)
	}

	bfSearch(START)

	for _, v := range bfTree { //探索木を表示する
		fmt.Printf("(%d, %d)", v.startNode, v.endNode)
	}
	fmt.Println()
}
