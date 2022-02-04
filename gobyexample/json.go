package main

import (
	"encoding/json"
	// "os"
	"fmt"
)

type testStruct struct {
	Int1 int
	Float1 float64
	Intarr []int
}

func testMarshal(v interface{}) {
	B, _ := json.Marshal(v)
	fmt.Println(string(B), B)
}

func main() {
	testMarshal(true)
	testMarshal(123)
	testMarshal(12.3)
	testMarshal('c')
	testMarshal("hello, world!")
	testMarshal([]int{1,2,3})
	testMarshal(map[int]int{1:50,2:100,3:150})

	st1 := &testStruct{
		Int1: 1,
		Float1: 2.0,
		Intarr: []int{3,4,5},
	}
	testMarshal(st1)

	byte1 := []byte(`{"Int1":1,"Float1":2,"Intarr":[3,4,5]}`)
	res := testStruct{}
	json.Unmarshal(byte1, &res)
	fmt.Println(res)
	byte2 := []byte(`{"iNT1":1,"floaT1":2,"iNTArr":[3,4,5]}`)
	res2 := testStruct{}
	json.Unmarshal(byte2, &res2)
	fmt.Println(res2)
}