package main

import (
	"errors"
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
		if runes[consumed] != '(' {
			break
		}
		consumed++
		c, err := parse(runes[consumed:])
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

func parse(runes []rune) (int, error) {
	max := len(runes)
	if max == 0 {
		return 0, errors.New("Missing )")
	}
	consumed := 0
	for consumed < max {
		r := runes[consumed]
		consumed++
		switch r {
		case '(':
			c, e := parse(runes[consumed:])
			if e != nil {
				return consumed + c, e
			}
			consumed += c
		case ')':
			return consumed, nil
		}
	}
	return 0, errors.New("Missing )")
}
