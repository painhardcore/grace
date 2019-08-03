package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker() {
	fmt.Println("Start Working")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func main() {
	// catch interrupt signal
	signalCatcher := make(chan os.Signal)
	signal.Notify(signalCatcher, syscall.SIGINT)
	// signal when worker function is finished
	done := make(chan struct{})

	fmt.Println("Hello world")
	go func() {
		worker()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Test over")
}
