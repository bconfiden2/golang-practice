package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome("로꾸꺼로꾸꺼"))
	fmt.Println(isPalindrome("수박이박수"))
	fmt.Println(isPalindrome("여보게저기저게보여"))
	fmt.Println(isPalindrome("다시합창합시다"))
	fmt.Println(isPalindrome("hello, world!"))
}

func isPalindrome(str string) bool {
	target := []rune(str)
	for i := 0 ; i < len(target) / 2 ; i++ {
		j := len(target) - i - 1
		if target[i] != target[j] {
			return false
		}
	}
	return true
}