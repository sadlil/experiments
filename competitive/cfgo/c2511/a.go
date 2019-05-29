package main

import "fmt"

func main() {
	var number int64
	fmt.Scanf("%d", &number)

	var val int64 = 1
	number = number - 2

	mod := (number) % 3
	if mod == 0 {
		val++
		number = number - 1
	}
	fmt.Printf("%d %d %d", 1, val, number)
}
