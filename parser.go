package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	runes := []rune(strings.TrimSpace(os.Args[1]))
	l := len(runes)
	consumed := 0
	for consumed < l {
		c, err := parse(runes[consumed:], 0)
		if err != nil {
			log.Fatalf("not balanced: %v", err)
		}
		consumed += c
	}

	if consumed != l {
		log.Fatal(fmt.Errorf("Not balanced: read %d of %d characters", consumed, l))
	}
	fmt.Printf("balanced\n")
}

func parse(runes []rune, depth int) (int, error) {
	if runes[0] != '(' {
		return 0, fmt.Errorf("depth %d, substring %q doesn't begin with (", depth, string(runes))
	}
	if len(runes) < 2 {
		return 0, fmt.Errorf("depth %d, substring %q missing )", depth, string(runes))
	}
	consumed := 1
	max := len(runes)
	for r := runes[consumed]; consumed < max; r = runes[consumed] {
		fmt.Printf("depth %d, consumed %d, '%c', %q\n", depth, consumed, r, string(runes[consumed:]))
		switch r {
		case '(':
			if c, e := parse(runes[consumed:], depth+1); e != nil {
				return consumed + c, e
			} else {
				consumed += c
			}
		case ')':
			consumed++
			fmt.Printf("depth %d, returning %d, nil\n", depth, consumed)
			return consumed, nil
		}
	}
	return consumed, fmt.Errorf("depth %d, missing matching ), consumed %d of %d", depth, consumed, max)
}
