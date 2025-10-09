package main

/*
 * Recognized a Dyck language by destructively reducing
 * pairs of matching parens/braces/brackets to the zero-length
 * string.
 */

import (
	"flag"
	"fmt"
)

func main() {
	matchPairs := flag.String("m", "(){}[]", "pairs of matching characters")
	flag.Parse()

	matches, parens := matchesAndParens(*matchPairs)

	runes := []rune(flag.Arg(0))

	i := 0
	for {
		fmt.Printf("%d String: %q\n", i, string(runes))
		if len(runes) < 2 {
			// one remaning unmatched paren/brace/bracket
			break
		}
		if i == len(runes)-1 {
			// got to the end of the slice without finding
			// an opening paren/brace/bracket
			break
		}
		if parens[runes[i]] {
			if thesis, ok := matches[runes[i]]; ok {
				if runes[i+1] == thesis {
					runes = append(runes[:i], runes[i+2:]...)
					i = 0 // start over at the beginning of runes[]
					continue
				}
			}
		}
		i++
	}

	if len(runes) != 0 {
		fmt.Printf("Expression unbalanced\n")
		return
	}
	fmt.Printf("Expression balanced\n")
}

func matchesAndParens(matchPairs string) (map[rune]rune, map[rune]bool) {
	matches := make(map[rune]rune)
	parens := make(map[rune]bool)
	var last rune
	for i, r := range []rune(matchPairs) {
		if i&0x01 == 1 {
			matches[last] = r // Find a ']', must find '[' on stack top
			continue
		}
		parens[r] = true
		last = r
	}
	return matches, parens
}
