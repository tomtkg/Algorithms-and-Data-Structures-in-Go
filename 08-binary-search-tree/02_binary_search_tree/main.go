package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type fruit struct {
	name  string
	price int
}

type node struct { //フルーツをノードとする二分木
	value  *fruit
	childL *node
	childR *node
}

func main() {
	var root *node
	var names []string
	_ = json.NewDecoder(os.Stdin).Decode(&names)

	// 各ノードを生成する
	data := [15]*node{
		{value: &fruit{"kiwi", 150}},
		{value: &fruit{"orange", 80}},
		{value: &fruit{"blueberry", 320}},
		{value: &fruit{"litchi", 50}},
		{value: &fruit{"papaya", 200}},
		{value: &fruit{"persimmon", 100}},
		{value: &fruit{"melon", 780}},
		{value: &fruit{"banana", 30}},
		{value: &fruit{"mango", 180}},
		{value: &fruit{"raspberry", 350}},
		{value: &fruit{"pineapple", 270}},
		{value: &fruit{"lemon", 70}},
		{value: &fruit{"apple", 120}},
		{value: &fruit{"durian", 980}},
		{value: &fruit{"pear", 90}},
	}

	// 二分探索木を構成する
	for _, v := range data {
		addNode(&root, v)
	}

	// できた木を表示する
	fmt.Println("-- Generated binary search tree --")
	printTree(root, 0)
	fmt.Println()

	// 果物の探索
	for _, name := range names {
		result := searchNode(root, name)
		if result == nil {
			fmt.Printf("%s does not exists\n", name)
		} else {
			fmt.Printf("%s: %d yen\n", result.name, result.price)
		}
	}
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

func searchNode(tree *node, name string) *fruit {
	if tree == nil {
		return nil
	}
	if tree.value.name == name {
		return tree.value
	}
	if tree.value.name > name {
		return searchNode(tree.childL, name)
	} else {
		return searchNode(tree.childR, name)
	}
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
