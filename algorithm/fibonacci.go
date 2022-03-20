package main

import "fmt"

func Fibonacci(n uint64) uint64 {
	if n == 0 {
		return 0
	}

	var n1, n2 uint64 = 0, 1
	for i := uint64(1) ; i < n ; i++ {
		n3 := n1 + n2
		n1, n2 = n2, n3
	}
	return n2
}

func main() {
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(Fibonacci(uint64(i)))
	}
}