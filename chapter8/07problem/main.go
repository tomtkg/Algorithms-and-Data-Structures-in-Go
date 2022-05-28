package main

import (
	"fmt"
)

var (
	stack []int // スタック用配列
	sp    int   // スタックポインタ
)

type node struct {
	num         int   // ノードの値
	left, right *node // 左右の部分木
}

func main() {
	p := &node{42,
		&node{15, nil,
			&node{18, nil,
				&node{28, nil, nil}}},
		&node{76,
			&node{63, nil, nil},
			&node{98,
				&node{92, nil, nil}, nil}}}

	// できた木を表示する
	fmt.Println("-- Binary tree --")
	printTree(p, 0)
	fmt.Println()

	// 探索
	fmt.Printf("(a) 63:")
	searchNode(p, 63)
	fmt.Println()

	// 行きがけ順トラバーサル
	fmt.Println("-- preorder trversal --")
	preorder(p)
	fmt.Println()

	// 返りがけ順トラバーサル
	fmt.Println("-- postorder trversal --")
	postorder(p)
	fmt.Println()

	// ノードの追加
	addNode(&p, &node{num: 30})

	// できた木を表示する
	fmt.Println("-- Generated binary tree --")
	printTree(p, 0)
	fmt.Println()
}

func printTree(tree *node, depth int) {
	if tree == nil {
		return
	}
	printTree(tree.left, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Printf("	")
	}
	fmt.Println(tree.num)
	printTree(tree.right, depth+1)
}

func searchNode(tree *node, num int) {
	if tree == nil {
		return
	}
	fmt.Printf(" -> %d", tree.num)
	if tree.num == num {
		return
	}
	if tree.num > num {
		searchNode(tree.left, num)
	} else {
		searchNode(tree.right, num)
	}
}

func preorder(tree *node) {
	if tree == nil {
		return
	}
	fmt.Printf(" -> %d", tree.num)
	preorder(tree.left)
	preorder(tree.right)
}

func postorder(tree *node) {
	if tree == nil {
		return
	}
	postorder(tree.left)
	postorder(tree.right)
	fmt.Printf(" -> %d", tree.num)
}

func addNode(tree **node, newNode *node) {
	if *tree == nil {
		*tree = newNode
		return
	}
	if (*tree).num > newNode.num {
		addNode(&(*tree).left, newNode)
	} else {
		addNode(&(*tree).right, newNode)
	}
}
