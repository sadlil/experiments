package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/mikespook/golib/signal"
)

func init() {
	signal.Bind(syscall.SIGTERM, func() uint {
		fmt.Println("called SIGTERM")
		time.Sleep(time.Minute * 2)
		fmt.Println("finishing SIGTERM")
		return signal.BreakExit
	})

	signal.Bind(syscall.SIGINT, func() uint {
		fmt.Println("called SIGINT")
		time.Sleep(time.Minute * 2)
		fmt.Println("finishing SIGINT")
		return signal.BreakExit
	})
}

func main() {
	fmt.Println("Application Started")
	go func() {
		for {
			fmt.Println("Running")
			time.Sleep(time.Second * 10)
		}
	}()

	signal.Wait()
}
