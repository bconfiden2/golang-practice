package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fibonacci(n-1) + fibonacci(n-2)
		}
	}
	fmt.Println(factorial(4))
	fmt.Println(fibonacci(10))
}