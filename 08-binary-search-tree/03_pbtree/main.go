package main

import "fmt"

var stack []int //スタック用配列
var sp int      //スタックポインタ

type node struct {
	num         int   //ノードの値
	left, right *node //左右の部分木
}

func pbtree(n int) *node {
	if n == 0 {
		return nil //ノード数が0なら，木は作らない
	}
	nLeft := (n - 1) / 2      //半分（切り捨て）は左部分木へ
	nRight := (n - 1) - nLeft //残りは右部分木へ
	p := &node{}              //ノードを作る
	p.left = pbtree(nLeft)    //左部分木を作る
	p.num = stack[sp]         //ノードにデータを読み込む
	sp++                      //スタックポインタを移動
	p.right = pbtree(nRight)  //右部分木を作る
	return p                  //できあがった木を返す
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

func main() {
	a() //図8.8
	b() //図8.9
	c() //章末(1)
	d() //章末(2)
}

func a() {
	fmt.Printf("図8.8")
	p := &node{0, nil,
		&node{10, nil,
			&node{20, nil,
				&node{30, nil,
					&node{40, nil,
						nil}}}}}

	// できた木を表示する
	fmt.Println("\n-- Binary tree --")
	printTree(p, 0)

	// 通りがけ順トラバーサル
	fmt.Println("\n-- inorder trversal --")
	inorder(p)

	// 完全にバランスした二分木を構成する
	p = pbtree(len(stack))

	// できた木を表示する
	fmt.Println("\n-- Generated binary tree --")
	printTree(p, 0)
}

func b() {
	fmt.Printf("図8.9")
	stack = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// 完全にバランスした二分木を構成する
	p := pbtree(len(stack))

	// できた木を表示する
	fmt.Println("\n-- Generated binary tree --")
	printTree(p, 0)
}

func c() {
	fmt.Printf("章末(1)")
	p := &node{42,
		&node{15, nil,
			&node{18, nil,
				&node{28, nil, nil}}},
		&node{76,
			&node{63, nil, nil},
			&node{98,
				&node{92, nil, nil}, nil}}}

	// できた木を表示する
	fmt.Println("\n-- Binary tree --")
	printTree(p, 0)

	fmt.Println("\n(a)")
	// 探索
	searchNode(p, 63)

	fmt.Printf("\n(b)")
	// 行きがけ順トラバーサル
	fmt.Println("\n-- preorder trversal --")
	preorder(p)
	// 返りがけ順トラバーサル
	fmt.Println("\n-- postorder trversal --")
	postorder(p)

	fmt.Printf("\n(c)")
	// ノードの追加
	addNode(&p, &node{num: 30})
	// できた木を表示する
	fmt.Println("\n-- Generated binary tree --")
	printTree(p, 0)
}

func d() {
	fmt.Printf("章末(2)")
	stack = []int{2, 9, 11, 28, 31, 45}

	// 完全にバランスした二分木を構成する
	p := pbtree(len(stack))

	// できた木を表示する
	fmt.Println("\n-- Generated binary tree --")
	printTree(p, 0)
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

func inorder(tree *node) {
	if tree == nil {
		return
	}
	inorder(tree.left)
	fmt.Println(tree.num)
	stack = append(stack, tree.num)
	inorder(tree.right)
}

func postorder(tree *node) {
	if tree == nil {
		return
	}
	postorder(tree.left)
	postorder(tree.right)
	fmt.Printf(" -> %d", tree.num)
}
