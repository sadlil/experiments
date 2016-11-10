package main

import (
	"fmt"
)

var value interface{}

func setValue(x int) {
	value = x
}

func setVаlue(x string) {
	value = x
}

func setVal𝚞e(x int, y string) {
	value = fmt.Sprintf("%d %s", x, y)
}

func main() {
	setValue(5)
	fmt.Println(value)

	setVаlue("five")
	fmt.Println(value)

	setVal𝚞e(5, "six")
	fmt.Println(value)
}
