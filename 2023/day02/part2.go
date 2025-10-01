package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input2.txt
var inputDay string

func main() {
	var total = 0
	for _, game := range strings.Split(inputDay, "\n") {
		marg := map[string]int{"green": 0, "red": 0, "blue": 0}
		for _, pull := range strings.Split(strings.Split(game, ":")[1], ";") {
			for _, color := range strings.Split(pull, ",") {
				amount, col := 0, ""
				fmt.Sscanf(color, "%d %s", &amount, &col)
				if amount >= marg[col] {
					marg[col] = amount
				}
			}
		}
		total += marg["green"] * marg["red"] * marg["blue"]
	}
	fmt.Println(total)
}
