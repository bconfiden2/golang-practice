package main

import "fmt"

func bubbleSort(arr []int) []int {
	for flg := true ; flg ; {
		flg = false
		for i := 0 ; i < len(arr)-1 ; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flg = true
			}
		}
	}
	return arr
}

func main() {
	arr := []int{10, 2, 1, 4, 6, 7, 9, 3, 5, 8}
	arr = bubbleSort(arr)
	fmt.Println(arr)
}