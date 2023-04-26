package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	// Find the shortest string in the array
	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	// Binary search for the longest common prefix
	low, high := 0, minLen-1
	for low <= high {
		mid := (low + high) / 2
		prefix := strs[0][:mid+1]

		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				high = mid - 1
				goto nextIteration
			}
		}

		low = mid + 1
		continue

	nextIteration:
	}

	return strs[0][:high+1]
}

func main() {
	// fl
	s1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s1))

	// ""
	s2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(s2))
}
