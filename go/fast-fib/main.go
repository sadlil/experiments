package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(fib(i))
	}
}

func fib(n int) int64 {
	if n < 0 {
		panic("number is negetive")
	}
	a, _ := fibTuples(n)
	return a
}

func fibTuples(n int) (int64, int64) {
	if n == 0 {
		return int64(0), int64(1)
	}
	a, b := fibTuples(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b

	if n%2 == 0 {
		return c, d
	} else {
		return d, c + d
	}
}
