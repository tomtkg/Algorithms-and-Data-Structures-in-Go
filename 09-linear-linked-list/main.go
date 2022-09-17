package main

import "fmt"

type cell struct { //セル
	value *int  //セルの値
	next  *cell //次のセルへのポインタ
}

func insertCell(p *cell, num int) {
	c := p
	for ; c.next.value != nil; c = c.next {
	}
	c.next = &cell{&num, c.next}
}

func insertOrderCell(p *cell, num int) {
	c := p
	for ; c.next.value != nil; c = c.next {
		if num < *c.next.value {
			c.next = &cell{&num, c.next}
			return
		}
	}
	c.next = &cell{&num, c.next}
}

func deleteCell(p *cell, num int) {
	for c := p; c.next.value != nil; c = c.next {
		if *c.next.value == num {
			*c.next = *c.next.next
			return
		}
	}
}

func deleteEvenCell(p *cell) {
	for c := p; c.next.value != nil; {
		if *c.next.value%2 == 0 {
			*c.next = *c.next.next
			continue
		}
		c = c.next
	}
}

func printCell(p *cell) {
	for c := p.next; c.value != nil; c = c.next {
		fmt.Printf("%d,", *c.value)
	}
	fmt.Println()
}

func printMemory(p *cell) {
	fmt.Println(p, nil, p.next)
	for c := p.next; c.value != nil; c = c.next {
		fmt.Println(c, *c.value, c.next)
	}
}

func main() {
	a() //問9.1
	b() //問9.2
	c() //章末問題
}

func a() {
	fmt.Println("問9.1")
	p := &cell{nil, &cell{}}
	for _, num := range []int{1, 6, 10, 15, 21, 39, 44, 52, 71, 89, 93} {
		insertCell(p, num)
	}
	printCell(p)
	insertOrderCell(p, 9)
	printCell(p)
}

func b() {
	fmt.Println("問9.2")
	p := &cell{nil, &cell{}}
	for _, num := range []int{4, 1, 7, 9, 0, 6, 3, 5, 8, 2} {
		insertOrderCell(p, num)
	}
	printCell(p)
	deleteEvenCell(p)
	printCell(p)
}

func c() {
	fmt.Println("章末問題")
	p := &cell{nil, &cell{}}
	for _, num := range []int{84, 95, 78, 27, 56} {
		insertOrderCell(p, num)
	}
	printMemory(p)
	deleteCell(p, 78)
	printMemory(p)
}
