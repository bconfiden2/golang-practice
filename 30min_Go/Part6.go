package main

import (
	"fmt"
	"time"
)

var (
	A int = 0
	B, C int = 1, 2
)

func cum(a, b, c int, s string) {
	for i:=0 ; i < 5 ; i++ {
		A += a
		B += b
		C += c
		fmt.Println(s, A, B, C)
	}
}

func goroutine_test() {
	cum(1, 1, 1, "ori")
	go cum(100, 100, 100, "go1")
	go cum(10, 10, 10, "go2")
	

	time.Sleep(time.Millisecond)
}

func sum(arr []int, c chan int) {
	fmt.Printf("%p %p\n", &arr, c)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
	c <- sum+sum
	c <- sum*3
}

func channel_test() {
	var arr1 []int
	for i:=0 ; i < 5 ; i++ {
		arr1 = append(arr1, i+1)
	}
	arr2 := []int{10,20,30,40,50}

	c := make(chan int)
	fmt.Printf("%p %p %p\n", &arr1, &arr2, c)

	arr3 := []int{100,200,300}
	go sum(arr3, c)
	go sum(arr2, c)
	go sum(arr1, c)
	x := <- c
	y := <- c
	z := <-c

	fmt.Println(x,y,z)

	ch := make(chan int)
	go func() {ch <- 1} ()
	fmt.Println(<-ch)
	fmt.Println("-------")
	
	ch2 := make(chan int, 1)
	
	go func() {
		kk := <-ch2
		fmt.Println("go!", kk)
		ch2 <- kk * 10
	} ()
	ch2 <- 1
	fmt.Println(<-ch2)

	// time.Sleep(time.Second)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0 ; i < n ; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func range_close() {
	c := make(chan int)
	go fibonacci(16, c)
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	// goroutine_test()
	// channel_test()
	range_close()
}