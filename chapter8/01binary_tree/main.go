package main

import (
	"fmt"
)

type person struct {
	name string
	year int
}

type node struct {
	value  *person
	childL *node
	childR *node
}

func main() {
	// 各ノードを生成する
	gary := &node{value: &person{"Gary", 1997}}
	steven := &node{value: &person{"Steven", 1994}}
	tyler := &node{value: &person{"Tyler", 1985}}
	upton := &node{value: &person{"Upton", 1988}}
	martin := &node{value: &person{"Martin", 1973}}

	// ノードを二分木の形につなぐ
	tyler.childL = upton
	tyler.childR = martin
	gary.childL = steven
	gary.childR = tyler

	// できた木を表示する
	fmt.Println("-- Binary tree --")
	printTree(gary, 0)
}

// 二分木を表示する関数．各ノードは，深さに合わせて字下げする．
func printTree(tree *node, depth int) {
	if tree == nil {
		return
	}
	printTree(tree.childL, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Printf("	")
	}
	fmt.Println(*tree.value)
	printTree(tree.childR, depth+1)
}
