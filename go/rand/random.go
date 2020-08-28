package main

import (
	"fmt"
	mrand "math/rand"
	"sort"
	"time"
)

func main() {
	mrand.Seed(time.Now().Unix())

	mp := map[int]struct{}{}
	var result []int
	for i := 1; i <= 5; i++ {
		val := mrand.Intn(50)+1
		if _, ok := mp[val]; ok {
			i--
			continue
		}
		mp[val] = struct{}{}
		result = append(result, val)
	}

	sort.Ints(result)
	fmt.Printf("%v", result)
}
