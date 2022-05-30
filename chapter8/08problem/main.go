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
	stack = []int{2, 9, 11, 28, 31, 45}

	// 完全にバランスした二分木を構成する
	p := pbtree(len(stack))

	// できた木を表示する
	fmt.Println("-- Generated binary tree --")
	printTree(p, 0)
	fmt.Println()
}

func pbtree(n int) *node {
	if n == 0 {
		return nil
	}
	nLeft := (n - 1) / 2
	nRight := (n - 1) - nLeft
	p := &node{}
	p.left = pbtree(nLeft)
	p.num = stack[sp]
	sp++
	p.right = pbtree(nRight)
	return p
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
