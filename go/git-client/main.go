package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v3"
	"io"
)

func main() {
	r, err := git.NewRepository("https://github.com/sadlil/gologger", nil)
	if err != nil {
		panic(err)
	}

	if err := r.PullDefault(); err != nil {
		panic(err)
	}

	iter, err := r.Commits()
	if err != nil {
		panic(err)
	}
	defer iter.Close()

	for {
		//the commits are not sorted in any special order
		commit, err := iter.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		fmt.Println(commit, commit.Message)
	}
}
