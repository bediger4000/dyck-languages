package main

/*
 * Given a string of round, curly, and square open and closing brackets,
 * return whether the brackets are balanced (well-formed).
 *
 * Based on Theorem 2 of "Language Recognition by Marking Automata",
 * R.W. Ritchie, F. N. Springsteel, in "Information and Control" 30,
 * 313-300 (1972).
 */

import (
	"fmt"
	"os"
)

var matching = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func main() {
	str := []rune(os.Args[1])

	depth, valid := calculateDepth(str)
	if !valid {
		fmt.Printf("unbalanced expression\n")
		return
	}

	mismatch := false

foundMismatch:
	for d := 1; d <= depth; d++ {
		currentDepth := 0
		var thesis rune
		for _, r := range str {
			switch r {
			case '(', '[', '{':
				currentDepth++
				if currentDepth == d {
					thesis = matching[r]
				}
			case ')', ']', '}':
				if currentDepth == d {
					if r != thesis {
						mismatch = true
						break foundMismatch
					}
				}
				currentDepth--
			}
		}
	}

	if mismatch {
		fmt.Printf("expression unbalanced\n")
		return
	}
	fmt.Printf("expression balanced\n")
}

// calculateDepth finds the max "depth" of nesting,
// without regard to matching. Returns depth and true
// if there's an even number of parens, and there's a
// closing paren/brace/bracket for every opening paren/brace/bracket
func calculateDepth(str []rune) (int, bool) {
	var depth, maxDepth int

unbalanced:
	for _, r := range str {
		switch r {
		case '(', '[', '{':
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		case ')', ']', '}':
			depth--
			if depth < 0 {
				break unbalanced
			}
		}
	}

	if depth != 0 {
		return 0, false
	}

	return maxDepth, true
}
