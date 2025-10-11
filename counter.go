package main

/*
 * Dyck-1 language recognizer, counter approach
 */

import (
	"fmt"
	"os"
)

func main() {
	var depth, maxDepth int

unbalanced:
	for _, r := range []rune(os.Args[1]) {
		switch r {
		case '(':
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		case ')':
			depth--
			if depth < 0 {
				break unbalanced
			}
		}
	}

	if depth != 0 {
		fmt.Printf("unbalanced expression\n")
		return
	}
	fmt.Printf("balanced expression, %d max depth\n", maxDepth)
}
