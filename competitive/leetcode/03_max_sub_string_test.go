package leetcode

import (
	"testing"
)

// Solution Approach - https://leetcode.com/problems/longest-substring-without-repeating-characters/solution/
//  1. Brute Force - O(n3) - Check all the substring one by one to see if it has no duplicate character.
//  2. Sliding Window with Set - O(2n) - a substring s_{ij} from index i to j−1 is already checked to
//  have no duplicate characters. We only need to check if s[j] is already in the substring s_{ij}.
//  3. Sliding Window with HashMap - O(n) - store the position of the character occurrence.

// the following solution applies the 3rd approach for the problem.
func lengthOfLongestSubstring(s string) int {
	var indexMp = map[rune]int{}

	var result int

	var ri int
	for idx, r := range s {
		var dist int
		pos, ok := indexMp[r]
		if ok {
			if pos > ri {
				ri = pos
			}
		}

		dist = idx - ri + 1
		indexMp[r] = idx + 1

		if dist > result {
			result = dist
		}
	}

	return result
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct{
		input string
		output int
	}{
		{
			input: "c",
			output: 1,
		},
		{
			input: "",
			output: 0,
		},
		{
			input: " ",
			output: 1,
		},
		{
			input: "abba",
			output: 2,
		},
		{
			input: "abcabc",
			output: 3,
		},
		{
			input: "abcabcbb",
			output: 3,
		},
		{
			input: "pwwkew",
			output: 3,
		},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if test.output != result {
			t.Error("Expected", test.output, "Got", result)
			t.Fail()
		}
	}
}