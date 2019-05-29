package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

	var input int64
	fmt.Fscanf(scanner, "%d\n", &input)

	var maxX int64 = -1//; var maxY int64 = -1
	for i := int64(0); i < input; i++ {
		var x, y int64
		fmt.Fscanf(scanner, "%d %d\n", &x, &y)
		//// fmt.Printf("%d %d %d/%d", x, y, i, input)
		//if x == y {
		//	maxX = maxI(maxX, x)
		//} else {
		//	maxY = maxI(maxY, maxI(x, y))
		//}
		maxX = maxI(maxX, x+y)
	}
	fmt.Fprintf(scanner, "%d", maxX)

	// fmt.Printf("max %d %d", maxX, maxY)
	//if maxY > maxX {
	//	fmt.Printf("%d\n", maxY+1)
	//} else {
	//	fmt.Printf("%d\n", maxX*2)
	//}
	scanner.Flush()
}

func maxI(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
