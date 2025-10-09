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

	if !validate(str) {
		fmt.Printf("unbalanced expression\n")
		return
	}

	depth := calculateDepth(str)

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

func calculateDepth(str []rune) int {
	d, depth := 0, 0
	for _, r := range str {
		switch r {
		case '(', '[', '{':
			d++
			if d > depth {
				depth = d
			}
		case ')', ']', '}':
			d--
		}
	}
	return depth
}

func validate(str []rune) bool {
	if len(str)%2 != 0 {
		return false
	}

	var parens, braces, brackets [2]int

	for _, r := range str {
		switch r {
		case '(':
			parens[0]++
		case ')':
			parens[1]++
		case '[':
			brackets[0]++
		case ']':
			brackets[1]++
		case '{':
			braces[0]++
		case '}':
			braces[1]++
		}
	}

	if parens[0] != parens[1] {
		return false
	}
	if brackets[0] != brackets[1] {
		return false
	}
	if braces[0] != braces[1] {
		return false
	}

	return true
}
