package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go read(ch)
	go write(ch)
	go write(ch)

	select {}
}

func read(c chan string) {
	for val := range c {
		fmt.Println(val)
	}
}

func write(c chan string) {
	for {
		c <- time.Now().String()
		time.Sleep(time.Second * 10)
	}
}
