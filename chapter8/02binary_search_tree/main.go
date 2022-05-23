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
	data := &node{value: &person{"Minks", 1985},
		childL: &node{value: &person{"Cindy", 1983},
			childL: &node{value: &person{"Bridget", 1982}},
			childR: &node{value: &person{"Luther", 1989},
				childL: &node{value: &person{"Jack", 1995}},
				childR: &node{value: &person{"Maroon", 1998}},
			},
		},
		childR: &node{value: &person{"Rady", 1998},
			childL: &node{value: &person{"Minuet", 1978}},
			childR: &node{value: &person{"Ruby", 1995},
				childR: &node{value: &person{"Tiki", 1982}},
			},
		},
	}

	// できた木を表示する
	fmt.Println("-- Binary search tree --")
	printTree(data, 0)
	fmt.Println()

	// ノードの探索と追加
	searchNode(data, "Maroon")
	fmt.Println()
	addNode(&data, &node{value: &person{"Elpe", 1982}})

	// できた木を表示する
	fmt.Println("-- Binary search tree --")
	printTree(data, 0)
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

func searchNode(tree *node, name string) {
	if tree == nil {
		return
	}
	fmt.Println(*tree.value)
	if tree.value.name == name {
		return
	}
	if tree.value.name > name {
		searchNode(tree.childL, name)
	} else {
		searchNode(tree.childR, name)
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
