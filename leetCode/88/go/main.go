package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	return nums1
}

func main() {
	s1 := []int{1, 2, 3, 0, 0, 0}
	s2 := []int{2, 5, 6}
	fmt.Println(merge(s1, 3, s2, 3))
}
