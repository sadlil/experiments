package algo

import (
	"fmt"
	"testing"
)

func dp(n, i, sum int, ar []int) int {
	if i>=n {
		return sum
	}

	first := dp(n, i+1, sum, ar)
	second := dp(n, i+2, sum + ar[i], ar)

	if first > second {
		return first
	}
	return second
}

func TestContMain(t *testing.T) {
	var ar = []int{10, 1, 1, 100}
	fmt.Println(dp(len(ar), 0, 0, ar))

	ar = []int{10, 20, 1}
	fmt.Println(dp(len(ar), 0, 0, ar))
}