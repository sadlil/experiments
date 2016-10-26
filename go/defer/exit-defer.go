package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Main")
	defer func() {
		fmt.Println("Inside defer")
	}()

	fmt.Println("Ending pkg")
	os.Exit(0)
}
