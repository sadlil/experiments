package main

import "fmt"

func main() {
	fmt.Println("running without sleep")
	loop()
}

func loop() {
	for {
		var _ = 0
	}
}