package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile(".env")
	if err != nil {
		fmt.Println(err)
	}

	array := strings.Split(string(data), "\n")
	for _, a := range array {
		if strings.Contains(a, "=") {
			keyValuePair := strings.Split(a, "=")
			if len(keyValuePair) == 2 {
				os.Setenv(keyValuePair[0], keyValuePair[1])
			}
		}
	}

	vals := os.Environ()
	for _, v := range vals {
		fmt.Println(v, os.Getenv(v))
	}
}
