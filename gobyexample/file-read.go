package main

import (
	"os"
	"fmt"
	"bufio"
)

func readFromFile(f *os.File, sz int) {

	bt := make([]byte, sz)
	n, err := f.Read(bt)
	check(err)
	fmt.Printf("%d bytes: %s\n", n, string(bt))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, e := os.ReadFile("input")
	check(e)
	fmt.Println(string(data))

	f, e := os.Open("input")
	fmt.Printf("return type: %T\n", f)
	check(e)
	readFromFile(f, 3)
	readFromFile(f, 12)

	_, err := f.Seek(3, 0)
	check(err)
	readFromFile(f, 4)

	_, err = f.Seek(0, 0)
	check(err)
	r := bufio.NewReader(f)
	b, err := r.Peek(3)
	check(err)
	fmt.Println(string(b))
}