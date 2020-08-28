package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve() {
	var st int
	Readf("%d", &st)

	a := -0.195959
	b := 8.705033
	c := -14.626206

	ret :=  a*float64(st*st) + b*float64(st) + c
	Writef("%d\n", int(math.Ceil(ret)))
}

func main()  {
	// If io still contains some buffered data flush it before
	// quiting
	defer io.Flush()

	// Solution of the problem goes here
	solve()
}

var (
	// fmt.Print and fmt.Scan is slow, we will use bufio
	// for read and write from stdin.
	io = bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)
)

func Readf(format string, args ...interface{})  {
	fmt.Fscanf(io, format, args...)
}

func Writef(format string, args ...interface{})  {
	fmt.Fprintf(io, format, args...)
}
