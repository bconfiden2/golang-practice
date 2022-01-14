package main

import (
	"fmt"
	"reflect"
)

type Coord struct {
	X, Y int
}

var (
	c1 = Coord{1,2}
)

func Pointer_Struct() {
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

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func array_slice() {
	var a[3] string;
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "Golang"
	fmt.Println(a)
	fmt.Println(a[0], a[2])

	arr := [3] int{1,2,3}
	fmt.Println(arr)

	slc := arr[:]
	var slc2 []int;
	slc2 = slc[:2]
	fmt.Println(slc, slc2)
	slc2[1] = 1000
	fmt.Println("슬라이스에서 값 변경시", slc, slc2)
	arr[0] = -20
	fmt.Println("arr 를 바꿨을 때 슬라이스들", slc, slc2)


	fmt.Println("배열의 타입", reflect.TypeOf(arr))
	fmt.Println("슬라이스 타입", reflect.TypeOf(slc))
	fmt.Println("슬라이스에서 인덱스", reflect.TypeOf(slc[0]))

	q := []int{1,2,3,4,5,6,7}
	fmt.Println("슬라이스로 초기화시 배열을 만든 뒤 배열참조하는 슬라이스를 만듦")
	fmt.Println(reflect.TypeOf(q))

	var test []int;
	test2 := make([]int, 5)
	test3 := make([]int, 1, 5)
	fmt.Println(test, test2, test3)

	test2[1] = -1
	test2 = append(test2, 100)
	test2 = append(test2, 101)
	test2 = append(test2, 102,103,104,105,106)

	printSlice(test2)
	printSlice(test3)

	var st1 []int = test2[:3]
	printSlice(st1)
	st2 := test3[:5]
	printSlice(st2)
	var st3 []int;
	st3 = st1[1:]
	printSlice(st3)
	st3 = append(st3, 1,2,3)

	fmt.Println("st3 에 1,2,3 추가했을 때")
	printSlice(test2)
	printSlice(st1)
	printSlice(st3)
	fmtPrintln("슬라이스는 원래꺼에서 레퍼런스 걸리니까 주의할것")
}

func main() {
	// Pointer_Struct()
	array_slice()
}