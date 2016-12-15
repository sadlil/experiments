package main

import (
	"encoding/json"
	"fmt"
)

type JsonObj struct {
	Name     string
	FullName string `json:"Name"`
}

func main() {
	testData := `{
	"Name": "foo",
	"name": "bar"
}`

	out := &JsonObj{}
	err := json.Unmarshal([]byte(testData), out)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out.Name, "||", out.FullName)
}
