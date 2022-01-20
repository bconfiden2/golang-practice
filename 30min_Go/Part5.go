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

func interface_test() {
	var i1 I = &T{1, 2}
	fmt.Printf("%p\n", i1)
	i1.m1(3,4)
	fmt.Println(i1)
	i1.m2()

	describe(i1)

	var x X = X{1,2,3,"hello",4}
	fmt.Println(x)
	var i2 I = x
	i2.m1(3,4)
	fmt.Println(i2)
	

	var y interface{}
	y = 42
	describe_empty(y)
	y = 1.5
	describe_empty(y)
	y = "hello"
	describe_empty(y)
	y = true
	describe_empty(y)
}

func describe_empty(x interface{}) {
	fmt.Printf("%p, %v, %T\n", &x, x, x)
}

func describe(x interface{ m1(int,int); m2() }) {
	fmt.Printf("%p\n", x)
}

type I interface {
	m1(int, int)
	m2()
}

type X struct {
	a, b, c int
	s string
	d int
}

func (x X) m1(p int, q int) {
	x.a += p
	x.b += q
}

func (x X) m2() {

}

type T struct {
	a, b int
}

func (t *T) m1(a, b int) {
	fmt.Printf("%p\n", t)
	fmt.Println(a, b)
	t.a += a
	t.b *= b
}

func (t *T) m2() {
	fmt.Printf("%p\n", t)
}

func main() {
	method_receiver()
	interface_test()
}