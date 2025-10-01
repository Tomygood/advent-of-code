package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input1.txt
var inputDay string

func main() {
	var total = 0
	marg := map[string]int{"green": 13, "red": 12, "blue": 14}
	for k, game := range strings.Split(inputDay, "\n") {
		var valid = true
		for _, pull := range strings.Split(strings.Split(game, ":")[1], ";") {
			for _, color := range strings.Split(pull, ",") {
				amount, col := 0, ""
				fmt.Sscanf(color, "%d %s", &amount, &col)
				if amount > marg[col] {
					valid = false
				}
			}
		}
		if valid {
			total += k + 1
		}
	}
	fmt.Println(total)
}
