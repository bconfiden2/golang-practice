package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20
	fmt.Println(m, len(m))

	var v1 int
	v1 = m["key1"]
	fmt.Println(v1)
	
	delete(m, "key1")
	fmt.Println(m, len(m))

	a, b := m["key1"]
	fmt.Println(a, b)
	a, b = m["key2"]
	fmt.Println(a, b)

	m["key3"] = 30
	m["key4"] = 40

	for k,v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}
}