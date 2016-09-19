package main

import "plugin"

func main() {
	p, err := plugin.Open("test-plugin/test-plugin.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("PrintHello")
	if err != nil {
		panic(err)
	}
	f.(func())()
}
