package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func possible(des string, patterns map[string]struct{}, seen map[string]int) int {
	if seen[des] > 0 {
		return seen[des]
	}
	if des == "" {
		return 1
	}

	res := 0
	if _, ok := patterns[des]; ok {
		res++
	}
	for st := range des {
		if _, ok := patterns[des[:st]]; ok {
			res += possible(des[st:], patterns, seen)
		}
	}
	seen[des] = res

	return res
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
	seen := make(map[string]int)

	for _, des := range designs {
		res += possible(des, patterns, seen)
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
