package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func contains(dict map[string]int, el string) int {
	for k := 3; k <= 5; k++ {
		if len(el) > k {
			val, ok := dict[el[:k]]
			if ok {
				return val
			}
		}
	}
	return 0
}

func main() {
	var total = 0
	var dict = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for _, line := range strings.Split(inputDay, "\n") {
		found := false
		dig := 0
		calib := 0
		for k, c := range line {
			let := contains(dict, line[k:])
			if c-48 >= 0 && c-48 <= 9 {
				if !found {
					calib = int(c - 48)
					found = true
				}
				dig = int(c - 48)
			}
			if let != 0 {
				if !found {
					calib = let
					found = true
				}
				dig = let
			}
		}
		calib = calib*10 + dig
		total += calib
	}
	fmt.Println(total)
}
