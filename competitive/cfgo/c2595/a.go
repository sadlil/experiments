// //+build a

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve() {
	var N int
	Readf("%d\n", &N)
	for i := 0; i < N; i++ {
		var q int
		Readf("%d\n", &q)

		var arr = make([]int, 0)
		for j := 0; j < q; j++ {
			var x int
			Readf("%d ", &x)
			arr = append(arr, x)
		}

		sort.Ints(arr)

		var count = 1
		for j := 1; j < q; j++ {
			if abs(arr[j-1]-arr[j]) == 1 {
				count++
				break
			}
		}
		Writef("%d\n", count)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
