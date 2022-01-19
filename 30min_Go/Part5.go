package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	fmt.Printf("포인터 아닐때 : %p\n", &v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) multiple(value int) {
	fmt.Printf("포인터일때 : %p\n", v)
	v.X = v.X * float64(value)
	v.Y = v.Y * float64(value)
	fmt.Println(v)
}

func method_receiver() {
	v := Vertex{3, 4}
	fmt.Printf("원래 주소 : %p\n", &v)
	fmt.Println(v.abs())
	fmt.Println(v)
	v.multiple(10)

	var p *Vertex
	p = &v
	fmt.Println(p)
	fmt.Printf("%p\n", p)
	p.multiple(10)
}

func main() {
	method_receiver()
}