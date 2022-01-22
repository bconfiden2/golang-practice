package main

import (
	"fmt"
	"time"
)

var (
	A int = 0
	B, C int = 1, 2
)

func cum(a, b, c int, s string) {
	for i:=0 ; i < 5 ; i++ {
		A += a
		B += b
		C += c
		fmt.Println(s, A, B, C)
	}
}

func goroutine_test() {
	cum(1, 1, 1, "ori")
	go cum(100, 100, 100, "go1")
	go cum(10, 10, 10, "go2")
	

	time.Sleep(time.Millisecond)
}

func main() {
	goroutine_test()
}