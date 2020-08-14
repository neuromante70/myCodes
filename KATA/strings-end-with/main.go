package main

// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).
// Examples:
// solution("abc", "bc") // returns true
// solution("abc", "d") // returns false

import (
	"fmt"
	s "strings"
)

func solution(str, ending string) bool {
	// Your code here!
	// str, substring = "abc", "bc"
	if s.HasSuffix(str, ending) {
		return true
	} else {
		return false
	}
}

func main() {
	result := solution("abc", "bc")
	fmt.Println(result)
}
