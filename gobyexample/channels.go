package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("고루틴 진입")
		var aa int
		aa = <-c
		fmt.Println("고루틴이 채널에서 데이터 수신 ", aa)
	} ()
	fmt.Println("메인프로세스")
	time.Sleep(time.Millisecond * 1)
	fmt.Println("메인프로세스 sleep 종료")
	c <- 2
	fmt.Println("메인프로세스가 채널에 데이터 전송완료")
	time.Sleep(time.Millisecond * 1)
	fmt.Println("메인프로세스 종료")
}