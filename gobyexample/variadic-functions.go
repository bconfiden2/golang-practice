package main

import "fmt"

func sum(nums ...int) (cnt, total int) {
	fmt.Printf("%p\n", &nums)
	cnt, total = 0, 0
	for _, v := range nums {
		cnt += 1
		total += v
	}
	return cnt, total
}

func main() {
	sum(1)
	sum(1,2,3)
	sum(1,2,3,4,5)
	
	nums := []int{1,2,3,4,5,6,7}
	fmt.Printf("%p\n", &nums)
	cnt, total := sum(nums...)
	fmt.Println(cnt, total)
}