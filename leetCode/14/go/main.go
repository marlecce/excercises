package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	// Binary search for the longest common prefix
	low, high := 0, len(strs[0])-1
	for low <= high {
		mid := (low + high) / 2
		prefix := strs[0][:mid+1]
		shortest := true

		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				shortest = false
				break
			}
		}

		if shortest {
			low = mid + 1
		} else {
			high = mid - 1
		}
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
