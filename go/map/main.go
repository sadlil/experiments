package main

import "fmt"

var mp = map[string]string{
	"foo": "bar",
	"ok":  "gotit",
}

func main() {
	fmt.Println(retMapOk())
	fmt.Println(retMapNotOk())
}

func retMapOk() string {
	fmt.Println(mp["foo"])
	a, b := mp["foo"]
	fmt.Println(a, b)
	return mp["foo"]
}

func retMapNotOk() string {
	fmt.Println(mp["no"])
	a, b := mp["no"]
	fmt.Println(a, b)
	return mp["no"]
}
