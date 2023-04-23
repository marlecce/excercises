package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack [10000]byte
	top := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if v, ok := m[c]; ok {
			stack[top] = v
			top++
		} else if top == 0 || c != stack[top-1] {
			return false
		} else {
			top--
		}
	}
	return top == 0
}

func main() {
	s1 := "()"
	fmt.Println(isValid(s1))

	s2 := "(]"
	fmt.Println(isValid(s2))
}
