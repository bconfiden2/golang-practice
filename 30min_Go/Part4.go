package main

import "fmt"

type Coord struct {
	X, Y int
}

var (
	c1 = Coord{1,2}
)

func main() {
	var p *int
	i, j := 10, 20
	
	p = &i
	q := &j
	fmt.Println("포인터 p 의 값 : ", p)
	fmt.Println("포인터 p 가 가리키는 값 :", *p)

	fmt.Println("*q로 바꾸기 전 :", j)
	*q = 500
	fmt.Println("*q로 바꾼 뒤 :", j)

	var t, u *int
	t, u = p, q
	fmt.Println(*t, *u)

	fmt.Println("초기화시 모든 값 입력", c1)
	c2 := Coord{Y:4}
	fmt.Println("초기화 시 Y에만 값 입력 ", c2)
	var c3 Coord = Coord{}
	fmt.Println("초기화 시 {}만", c3)
	c3.X = *p
	c3.Y = *u
	fmt.Println("포인터 p랑 u로 값 세팅 ", c3)

	pst := &c2
	pst.X = 1000
	fmt.Println("구조체 포인터로 값 세팅 ", c2)

	fmt.Println("구조체 포인터: ", pst)
	fmt.Println("해당 구조체 주소: ", &c2)
	fmt.Println("구조체 포인터가 가리키는 구조체: ", *pst)
}