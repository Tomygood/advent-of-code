package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func possible(des string, patterns map[string]struct{}, seen map[string]bool) bool {
	if _, ok := seen[des]; ok {
		return seen[des]
	}

	if des == "" {
		return true
	}

	if _, ok := patterns[des]; ok {
		return true
	}

	for st := range des {
		if _, ok := patterns[des[:st]]; ok {
			if possible(des[st:], patterns, seen) {
				seen[des] = true
				return true
			}
		}
	}
	seen[des] = false
	return false
}

func main() {
	start := time.Now()

	file := strings.Split(inputDay, "\r\n\r\n")

	patterns_l := strings.Split(file[0], ", ")
	patterns := make(map[string]struct{})
	for _, pat := range patterns_l {
		patterns[pat] = struct{}{}
	}

	designs := strings.Split(file[1], "\r\n")

	res := 0
	seen := make(map[string]bool)

	for _, des := range designs {
		if possible(des, patterns, seen) {
			res++
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
