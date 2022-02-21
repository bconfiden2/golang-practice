package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 2, 1, 4, 6, 7, 9, 3, 5, 8}
	arr = insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) ([]int) {
	for i := 1 ; i < len(arr) ; i++ {
		value := arr[i]
		var j int
		for j = i ; j > 0 && arr[j-1] >= value ; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = value
	}
	return arr
}