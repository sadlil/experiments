package main

import (
	"fmt"
	"time"
)

func goRoutine(num int, stop chan struct{}, writechan chan int) {
	for i := 1; i <= num; i++ {
		select {
		case <-stop:
			fmt.Println("stopping", num)
			return
		default:
		}
		fmt.Println("running", num)
		time.Sleep(time.Second * 5)
	}
	writechan <- num
}

func main() {
	stop := make(chan struct{})
	write := make(chan int)
	for i := 1; i <= 10; i++ {
		go goRoutine(i, stop, write)
	}

	res, ok := <-write
	if ok {
		fmt.Println("Finishes", res)
		close(stop)
		time.Sleep(time.Second * 2)
	}
}
