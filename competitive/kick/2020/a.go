package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	var T, N, K int
	Readf("%d", &T)

	for i := 1; i <= T; i++ {
		Readf("%d %d", &N, &K)

		var A int
		var count, expected int

		expected = K
		for j := 0; j < N; j++ {
			Readf("%d ", &A)
			if expected == 1 {
				count++
				expected = K
				continue
			}

			if expected == A {
				expected--
			}
		}
		Writef("Case #%d: %d\n", i, count)
	}
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
