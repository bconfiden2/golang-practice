package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		fmt.Println("\n", <-sigs)
		done <- true
	}()
	fmt.Println("waiting...")
	<-done
	fmt.Println("finish")
}