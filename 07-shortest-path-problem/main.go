package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	NODES       = 9  //ノードの数
	INF         = 99 //無限大の距離
	END_OF_PATH = -1 //経路の終端
)

type node struct {
	dist int //始点からの距離
	flag int //ノードが囲みの中にあるかどうかを表すフラグ
	path int //最短経路で目的のノードへ導く隣接ノード
}

var (
	START    int               //始点ノード
	END      int               //終点ノード
	node_dat [NODES]node       //ノード情報
	matrix   [NODES][NODES]int //隣接行列
)

func logPrint(num int) {
	log := [3][NODES]int{}
	for i, v := range node_dat {
		log[0][i] = v.dist
		log[1][i] = v.flag
		log[2][i] = v.path
	}
	fmt.Printf("node u = [%d]\n", num)
	fmt.Printf("dist = %2d\n", log[0])
	fmt.Printf("flag = %2d\n", log[1])
	fmt.Printf("path = %2d\n", log[2])
}

func dijkstra(start, end int) {
	//ノードstartを囲みの中に入れる
	node_dat[start].flag = 1

	//ノードstartに隣接するノードの設定
	for i, v := range matrix[start] {
		if v != INF {
			node_dat[i].dist = v
			node_dat[i].path = start
		}
	}

	//ノードstartの設定
	node_dat[start].dist = 0
	node_dat[start].path = END_OF_PATH

	logPrint(start)

	for u := start; u != end; {
		min := INF
		//囲みの中にないノードを調査する
		for i, v := range node_dat {
			if v.flag != 1 && min > v.dist {
				min = v.dist //最短経路長の最小値を求める
				u = i        //最小値をもつノードuを求める
			}
		}
		//最小値をもつノードuを囲みの中に入れる
		node_dat[u].flag = 1
		//囲みの中にないノードuの隣接ノードを調査する
		for i, v := range matrix[u] {
			if v != INF && node_dat[i].flag != 1 {
				if node_dat[i].dist > node_dat[u].dist+v { //ノードuを経由した方が距離が短いなら，
					node_dat[i].dist = node_dat[u].dist + v //距離を更新する
					node_dat[i].path = u                    //ノードuを経由するように最短経路を更新する
				}
			}
		}
		logPrint(u)
	}
}

func main() {
	graph := [][3]int{}
	_ = json.NewDecoder(os.Stdin).Decode(&graph)

	//初期化
	START, END = graph[0][0], graph[0][1]
	node_dat = [NODES]node{}
	for i := range node_dat {
		node_dat[i].dist = INF
	}
	matrix = [NODES][NODES]int{}
	for i := range matrix {
		for j := range matrix {
			matrix[i][j] = INF
		}
	}
	for _, v := range graph[1:] {
		matrix[v[0]][v[1]] = v[2]
		matrix[v[1]][v[0]] = v[2]
	}
	dijkstra(START, END) //ダイクストラ法による最短経路長探索

	//最短経路長探索および最短経路を表示する
	fmt.Printf("%d -> %d : Distance = %d, Path = %d", START, END, node_dat[END].dist, END)
	for u := END; node_dat[u].path != END_OF_PATH; u = node_dat[u].path {
		fmt.Printf(" %d", node_dat[u].path)
	}
	fmt.Println()
}
