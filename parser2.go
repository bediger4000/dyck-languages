package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	matchPairs := flag.String("m", "(){}[]", "pairs of matching characters")
	flag.Parse()

	setupMatches(*matchPairs)

	balanced(flag.Arg(0))
}

func balanced(str string) {
	runes := []rune(strings.TrimSpace(str))
	l := len(runes)
	consumed := 0
	for consumed < l {
		r := runes[consumed]
		consumed++
		closing, ok := matching[r]
		if !ok {
			log.Fatalf("No preceding opening character for closing character %c", r)
		}
		c, err := parse(runes[consumed:], closing)
		if err != nil {
			log.Fatalf("not balanced: %v", err)
		}
		consumed += c
	}

	if consumed != l {
		log.Fatalf("Not balanced: read %d of %d characters", consumed, l)
	}
	fmt.Printf("balanced\n")
}

func parse(runes []rune, closing rune) (int, error) {
	max := len(runes)
	if max == 0 {
		return 0, fmt.Errorf("Missing %c 1", closing)
	}
	consumed := 0
	for consumed < max {
		r := runes[consumed]
		consumed++
		if match, ok := matching[r]; ok {
			c, e := parse(runes[consumed:], match)
			if e != nil {
				return consumed + c, e
			}
			consumed += c
			continue
		}
		if r == closing {
			return consumed, nil
		}
		break
	}
	return 0, fmt.Errorf("Missing %c", closing)
}

var matching map[rune]rune

func setupMatches(matchPairs string) {
	matching = make(map[rune]rune)
	var last rune
	for i, r := range []rune(matchPairs) {
		if i&0x01 == 1 {
			matching[last] = r // look up '(' to find matching ')', or whatever pair
			continue
		}
		last = r
	}
}
