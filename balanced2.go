package main

/*
 * Determine whether strings of pairs of opening/closing characters
 * are "balanced". Allows arbitrary pairs, not just (,) [,] {,}
 */

import (
	"flag"
	"fmt"
)

func main() {

	matchPairs := flag.String("m", "(){}[]", "pairs of matching characters")

	flag.Parse()

	var stack []rune
	runes := []rune(flag.Arg(0))

	matches := setupMatches(*matchPairs)

	mismatch := false
	for _, r := range runes {
		if m, ok := matches[r]; len(stack) > 0 && ok {
			top := stack[len(stack)-1]
			if m != top {
				mismatch = true
				break
			}
			// pop matching opening "parentheses"
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, r)
	}

	if mismatch || len(stack) > 0 {
		// unmatched parens left on stack, or mismatched character
		fmt.Printf("Expression unbalanced\n")
		return
	}

	fmt.Printf("Expression balanced\n")
}

func setupMatches(matchPairs string) map[rune]rune {
	matches := make(map[rune]rune)
	var last rune
	i := 0
	for _, r := range matchPairs {
		switch i {
		case 0:
			last = r
			i = 1
		case 1:
			matches[r] = last // Find a ']', must find '[' on stack top
			i = 0
		}
	}
	return matches
}
