package main

import (
	"fmt"
)

func recoverControl() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in defer:\n", r)
	}
}

func main() {
	defer recoverControl()

	panic("panic!")
	panic("panic 2!")

	fmt.Println("hello, world!")
}