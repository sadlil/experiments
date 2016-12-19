package main

import (
	"fmt"
	"reflect"
)

type Foo struct{}

func main() {
	fmt.Println(reflect.TypeOf(&Foo{}).String())
	fmt.Println(reflect.TypeOf(Foo{}).String())
}
