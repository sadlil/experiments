package main

import "fmt"

func main() {
	fmt.Println("unning without sleep")
	loop()
}

func loop() {
	for {
		var _ = 0
	}
}