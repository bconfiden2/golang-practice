package main

import "fmt"

type embed struct {
	X, Y int
}

func (e *embed) test() int {
	fmt.Printf("%p\n", e)
	return e.X + e.Y
}

type container struct {
	*embed
	integer int
	str string
}

type sum interface {
	test() int
}

func main() {
	co := container{
		embed: &embed{
			X: 1,
			Y: 2,
		},
		integer: 4,
		str: "hi",
	}
	fmt.Println(co.X, co.Y)
	fmt.Println(co.embed.X, co.embed.Y)
	fmt.Println(co.test())
	fmt.Println(co.embed.test())

	var ifc sum = co.embed
	var ifc2 sum = &co
	fmt.Println(ifc.test())
	fmt.Println(ifc2.test())
}