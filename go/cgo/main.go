package main

/*
#include "hello.c"
#include <math.h>
*/
import "C"
import "fmt"

func main() {
	C.PrintHello()
	fmt.Println(C.sqrt(100))
}
