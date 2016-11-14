package main

import "fmt"

//easyjson:json
type JSONData struct {
	Data []string
}

func main() {
	d := &JSONData{}
	d.UnmarshalJSON([]byte(`{"Data" : ["One", "Two", "Three"]} `))
	fmt.Println(d)
}
