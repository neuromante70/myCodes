package main

import (
	"fmt"
	"strings"
)

func main() {
	//culo = ("%#v\n", strings.Split("abc", ""))
	// result := []string{strings.Split("abc", "")}
	runeSlice := strings.Split("abc", "")
	// runeSlice := []rune(cazzo)
	fmt.Printf("%#v\n", runeSlice)

	first := runeSlice[:1]
	fmt.Printf("First element: %v\n", first)
}
