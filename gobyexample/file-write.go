package main

import (
	"os"
	"bufio"
)

func main() {

	data := []byte("hello, world!\nbconfiden2\ntest\n")
	os.WriteFile("input", data, 0644)

	f, _ := os.Create("test.go")
	defer f.Close()

	data = []byte{97, 98, 99, 100}
	f.Write(data)
	f.WriteString("\nzzzzzzzzzzzzzzzzzzzzzzz\n")

	w := bufio.NewWriter(f)
	w.WriteString("bconfiden2\n")
	w.Flush()
}