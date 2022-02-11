package main

import (
	"fmt"
)

func main() {

	arr := make([]int, 10000)
	for i := 0 ; i < 10000 ; i++ {
		arr[i] = i
	}
	fmt.Println(binarySearch(arr, 3789))
	fmt.Println(binarySearch(arr, 6750))
}

func binarySearch(arr []int, target int) (mdx, cnt int) {
	ldx, rdx := 0, len(arr)-1
	for ldx <= rdx {
		mdx = (ldx + rdx) / 2
		if arr[mdx] < target {
			ldx = mdx + 1
		} else if arr[mdx] > target {
			rdx = mdx - 1
		} else {
			return mdx, cnt
		}
		cnt++
	}
	return -1, cnt
}