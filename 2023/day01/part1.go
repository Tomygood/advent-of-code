package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func main() {
	var total = 0
	for _, line := range strings.Split(inputDay, "\n") {
		found := false
		dig := 0
		calib := 0
		for _, c := range line {
			if c-48 >= 0 && c-48 <= 9 {
				if !found {
					calib = int(c - 48)
					found = true
				}
				dig = int(c - 48)
			}
		}
		calib = calib*10 + dig
		total += calib
	}
	fmt.Println(total)
}
