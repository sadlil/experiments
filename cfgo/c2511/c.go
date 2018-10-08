package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func gcd(a, b int64) int64  {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func solve() {
	var input int64
	Readf("%d\n", &input)

	var solution int64 = -1

	values := make([]int64, input)
	for i:=int64(0); i< input-1; i++ {
		Readf("%d ", &values[i])
	}
	Readf("%d\n", &values[input-1])

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	gcdAr := make([]int64, input+5)
	gcdAr[0] = values[0]
	for i:= int64(1); i<input; i++ {
		gcdAr[i] = gcd(values[i], gcdAr[i-1])
	}

	for i := input-1; i >= 0; i-- {
		if gcdAr[i] > values[i] {
			solution = i+1
		}
	}
	Writef("%d\n", solution)
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
