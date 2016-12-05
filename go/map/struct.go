package main

import "fmt"

type Struct struct {
	Field1 string
	Field2 string
}

func main() {
	mp := make(map[Struct]string)

	s1 := Struct{
		Field1: "foo",
	}
	mp[s1] = "foo"

	s2 := Struct{
		Field2: "bar",
	}
	mp[s2] = "bar"

	s3 := Struct{
		Field1: "foo",
		Field2: "bar",
	}
	mp[s3] = "foo-bar"

	for k, v := range mp {
		fmt.Println(k, v)
	}

	v, ok := mp[s1]
	fmt.Println(v, ok)

	v, ok = mp[s2]
	fmt.Println(v, ok)

	v, ok = mp[s3]
	fmt.Println(v, ok)
}
