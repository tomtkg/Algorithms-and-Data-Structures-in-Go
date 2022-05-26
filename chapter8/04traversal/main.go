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
	var root *node

	// 各ノードを生成する
	data := [10]*node{
		{value: &person{"da Vinci", 1452}},
		{value: &person{"Michelangelo", 1475}},
		{value: &person{"Raphael", 1483}},
		{value: &person{"Hokusai", 1760}},
		{value: &person{"van Gogh", 1853}},
		{value: &person{"Kandinsky", 1866}},
		{value: &person{"Matisse", 1869}},
		{value: &person{"Picasso", 1881}},
		{value: &person{"Utrillo", 1883}},
		{value: &person{"Dali", 1904}},
	}

	// 二分探索木を構成する．この例では，生年順に追加している．
	for _, v := range data {
		addNode(&root, v)
	}

	// できた木を表示する
	fmt.Println("-- Generated binary search tree --")
	printTree(root, 0)
	fmt.Println()

	// 行きがけ順トラバーサル
	fmt.Println("-- preorder trversal --")
	preorder(root)
	fmt.Println()

	// 通りがけ順トラバーサル
	fmt.Println("-- inorder trversal --")
	inorder(root)
	fmt.Println()

	// 返りがけ順トラバーサル
	fmt.Println("-- postorder trversal --")
	postorder(root)
	fmt.Println()
}

func addNode(tree **node, newNode *node) {
	if *tree == nil {
		*tree = newNode
		return
	}
	if (*tree).value.name > newNode.value.name {
		addNode(&(*tree).childL, newNode)
	} else {
		addNode(&(*tree).childR, newNode)
	}
}

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

func preorder(tree *node) {
	if tree == nil {
		return
	}
	fmt.Println(*tree.value)
	preorder(tree.childL)
	preorder(tree.childR)
}

func inorder(tree *node) {
	if tree == nil {
		return
	}
	inorder(tree.childL)
	fmt.Println(*tree.value)
	inorder(tree.childR)
}

func postorder(tree *node) {
	if tree == nil {
		return
	}
	postorder(tree.childL)
	postorder(tree.childR)
	fmt.Println(*tree.value)
}
