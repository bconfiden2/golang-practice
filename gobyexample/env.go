package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {

	os.Setenv("FOO", "test")
	fmt.Println(os.Getenv("FOO"))
	fmt.Println(os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(e, pair[0])
	}
}