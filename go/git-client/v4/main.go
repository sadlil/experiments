package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io"
)

func main() {
	r := git.NewMemoryRepository()
	options := &git.CloneOptions{
		URL: "git@github.com:sadlil/gologger.git",
	}

	if err := r.Clone(options); err != nil {
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
