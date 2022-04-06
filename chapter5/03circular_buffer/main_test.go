package main

import (
	"os"
	"testing"
)

func Test_enqueue(t *testing.T) {
	initQueue()
	for i := 0; i < QUEUESIZE; i++ {
		if got := enqueue(float64(i)); got != 1 {
			t.Errorf("%d: enqueue() = %v, want %v", i, got, 1)
		}
	}
	if got := enqueue(QUEUESIZE); got != 0 {
		t.Errorf("%d: enqueue() = %v, want %v", QUEUESIZE, got, 0)
	}
}
func Test_dequeueA(t *testing.T) {
	initQueue()
	var data float64
	if got := dequeue(&data); got != 0 {
		t.Errorf("dequeue() = %v, want %v", got, 0)
	}
}
func Test_dequeueB(t *testing.T) {
	initQueue()
	Test_enqueue(t)
	var data float64
	for i := 0; i < QUEUESIZE; i++ {
		if got := dequeue(&data); got != 1 {
			t.Errorf("%d: dequeue() = %v, want %v", i, got, 1)
		}
		if data != float64(i) {
			t.Errorf("data = %v, want %v", data, float64(i))
		}
	}
	if got := dequeue(&data); got != 0 {
		t.Errorf("dequeue() = %v, want %v", got, 0)
	}
}
func Example_main() {
	r, w, _ := os.Pipe()
	w.Write([]byte("[[3.14,1592,65],[3,5,-8,0.979],[3,5,-8,0.979]]"))
	w.Close()
	os.Stdin = r
	initQueue()
	main()
	// Output:
	// [0 0 0 0 0 0 0 0 0 0] 0 0 0
	// [3.14 1592 65 0 0 0 0 0 0 0] 0 3 3
	// 3.14
	// 1592
	// 65
	// [3.14 1592 65 0 0 0 0 0 0 0] 3 3 0
	// [3.14 1592 65 3 5 -8 0.979 0 0 0] 3 7 4
	// 3
	// 5
	// -8
	// 0.979
	// [3.14 1592 65 3 5 -8 0.979 0 0 0] 7 7 0
	// [0.979 1592 65 3 5 -8 0.979 3 5 -8] 7 1 4
	// 3
	// 5
	// -8
	// 0.979
	// [0.979 1592 65 3 5 -8 0.979 3 5 -8] 1 1 0
}
