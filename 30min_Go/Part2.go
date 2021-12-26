package main

import (
	"fmt"
	"math/cmplx"
)

var (
	i int
	f float64
	b bool
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	s string
)

func foo(x, y int) (int, int, float64) {
	var a, b int
	a = x + y
	b = x * y
	c := float64(x) / float64(y)
	return a, b, c
}

func main() {
	fmt.Printf("%T(%v)\n", MaxInt, MaxInt)
	fmt.Printf("%T(%v)\n", z, z)
	fmt.Println(foo(10, 20))
}