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
	dfFlag  [NODES]int        //探索状況（0: 未探索，1: 探索済）
	dfTree  [NODES - 1]edge   //探索木
	edgeCnt int
)

func dfSearch(u int) {
	dfFlag[u] = 1 //ノードuを探索済みにする

	for v := 0; v < NODES; v++ { //探索されていないノードvを調査する
		if matrix[u][v] == 1 && dfFlag[v] == 0 {
			dfTree[edgeCnt].startNode = u //辺(u, v)を探索木に登録する
			dfTree[edgeCnt].endNode = v
			edgeCnt++
			dfSearch(v) //ノードvを探索する
		}
	}
}
func main() {
	graph := [][2]int{}
	_ = json.NewDecoder(os.Stdin).Decode(&graph)

	//初期化
	matrix = [NODES][NODES]int{}
	dfFlag = [NODES]int{}
	edgeCnt = 0

	for _, v := range graph {
		matrix[v[0]][v[1]] = 1
		matrix[v[1]][v[0]] = 1
	}

	for _, v := range matrix { //隣接行列を表示する
		fmt.Println(v)
	}

	dfSearch(START)

	for _, v := range dfTree { //探索木を表示する
		fmt.Printf("(%d, %d)", v.startNode, v.endNode)
	}
	fmt.Println()
}
