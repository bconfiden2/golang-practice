package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend!")
	default:
		fmt.Println("weekday!")
	}

	myType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("boolean!")
		case int, int32, int64:
			fmt.Println("integer!")
		case float32, float64:
			fmt.Println("floating!")
		default:
			fmt.Println("etc!")
		}
	}
	myType(true)
	myType(10)
	myType(30.0)
	myType('a')
	myType("asdf")
}