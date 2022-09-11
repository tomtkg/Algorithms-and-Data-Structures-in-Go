package main

import (
	"mylib"
	"testing"
)

func Test_push(t *testing.T) {
	initStack()
	for i := 0; i < STACKSIZE; i++ {
		if got := push(float64(i)); got != 1 {
			t.Errorf("%d: push() = %v, want %v", i, got, 1)
		}
	}
	if got := push(STACKSIZE); got != 0 {
		t.Errorf("%d: enqueue() = %v, want %v", STACKSIZE, got, 0)
	}
}
func Test_popA(t *testing.T) {
	initStack()
	var data float64
	if got := pop(&data); got != 0 {
		t.Errorf("pop() = %v, want %v", got, 0)
	}
}
func Test_popB(t *testing.T) {
	initStack()
	Test_push(t)
	var data float64
	for i := 0; i < STACKSIZE; i++ {
		if got := pop(&data); got != 1 {
			t.Errorf("%d: pop() = %v, want %v", i, got, 1)
		}
		if data != float64(9-i) {
			t.Errorf("data = %v, want %v", data, float64(9-i))
		}
	}
	if got := pop(&data); got != 0 {
		t.Errorf("pop() = %v, want %v", got, 0)
	}
}
func Example_mainA() {
	mylib.SetString("1.2\n3.4\n1.6\n+\n*")
	main()
	// Output:
	// 6.00
}
func Example_mainB() {
	mylib.SetString("1.7\n2.8\n+\n2.5\n4.7\n-\n*\n-0.1\n+")
	main()
	// Output:
	// -10.00
}
