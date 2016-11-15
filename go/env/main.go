package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("NEW_ENV"))
	os.Setenv("NEW_ENV", "this is the value")
	fmt.Println(os.Getenv("NEW_ENV"))
}
