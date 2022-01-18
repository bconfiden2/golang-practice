package main

import "fmt"

func hanoi(n int, from, via, to string) {
	if n > 1 {
		hanoi(n-1, from, to, via)
		fmt.Println(from, " - ", to)
		hanoi(n-1, via, from, to)
	} else {
		fmt.Println(from, " - ", to)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	hanoi(n, "x", "y", "z")
}