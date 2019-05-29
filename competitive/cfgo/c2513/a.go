package main

import (
	"bufio"
	"os"
	"fmt"
)

func solve() {
	var n int64
	Readf("%d", &n)

	var st string
	Readf("%s", &st)

	//var flags = make([]int64, 12)
	for i := int64(0); i<n; i++ {
		val := st[i] - '0'
		Writef("%v", val)
	}
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
