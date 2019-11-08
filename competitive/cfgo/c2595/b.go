//+build b

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {

}

func main() {
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

func Readf(format string, args ...interface{}) {
	fmt.Fscanf(io, format, args...)
}

func Writef(format string, args ...interface{}) {
	fmt.Fprintf(io, format, args...)
}
