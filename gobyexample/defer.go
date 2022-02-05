package main

import (
	"fmt"
	"os"
)

func main() {
	defer removeFile()
	defer whenDefer(1)
	fmt.Println("before create")
	f := createFile("test.txt")
	fmt.Println("after create")
	defer closeFile(f)
	defer whenDefer(2)
	fmt.Println("before write")
	writeFile(f)
	fmt.Println("after write")
	fmt.Println("before main ends")
}

func removeFile() {
	fmt.Println("removeFile executed!")
	err := os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}

func createFile(path string) *os.File {
	fmt.Println("createFile executed!")
	if f, err := os.Create(path); err == nil {
		return f
	}
	return nil
}

func writeFile(f *os.File) {
	fmt.Println("writeFile executed!")
	fmt.Fprintln(f, "abcdefg")
	fmt.Fprintln(f, "testsets")
}

func closeFile(f *os.File) {
	fmt.Println("closeFile executed!")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func whenDefer(x int) {
	fmt.Println("Defer executed!", x)
}