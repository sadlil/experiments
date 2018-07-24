package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("running with sleep")
	loop()
}

func loop() {
	for {
		var _ = 0
		time.Sleep(time.Millisecond*500)
	}
}
