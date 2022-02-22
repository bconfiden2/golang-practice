package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 1, 6, 3, 2, 5, 8, 7, 9}
	fmt.Println(LIS(arr))
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func LIS(arr []int) (sz int) {
	dp := make([]int, len(arr))
	for i := range dp {
		dp[i] = 1
	}

	for i := range dp {
		for j := 0 ; j < i ; j++ {
			if arr[j] < arr[i] && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
				sz = Max(sz, dp[i])
			}
		}
	}
	return sz
}