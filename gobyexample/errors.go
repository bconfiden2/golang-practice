package main

import (
	"fmt"
	"errors"
	"strconv"
)

func err01(a int) (int, error) {
	if a >= 10 {
		return -1, errors.New(strconv.Itoa(a) + " is too high")
	}
	return a * 100, nil
}

type argError struct {
	arg int
	msg string
}

func (e* argError) Error() string {
	return fmt.Sprintf("%d %s", e.arg, e.msg)
}

func err02(a int) (int, error) {
	if a >= 10 {
		return -1, &argError{a, "is too high"}
	}
	return a * 100, nil
}

func main() {
	for _, v := range []int{3,8,11} {
		if ret, err := err01(v); err == nil {
			fmt.Println("success, ", ret)
		} else {
			fmt.Println("fail, ", err)
		}
	}

	for _, v := range []int{2,9,30} {
		if ret, err := err02(v); err == nil {
			fmt.Println("success, ", ret)
		} else {
			fmt.Println("fail, ", err)
		}
	}

	_, err := err02(50)
	fmt.Println()
	aE, flg := err.(*argError)
	fmt.Println(aE, flg)
}