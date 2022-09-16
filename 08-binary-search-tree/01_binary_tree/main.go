package main

import "fmt"

type person struct { //人の情報
	name string //名前
	year int    //誕生年
}

type node struct { //二分木
	value  *person //ノードの値
	childL *node   //左の子へのポインタ
	childR *node   //右の子へのポインタ
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

func main() {
	a() // 問8.1
	b() // 8.3節
	c() // 8.4節
}

func a() {
	fmt.Println("問8.1")
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

func b() {
	fmt.Println("8.3節")
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

func c() {
	fmt.Println("8.4節")
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
