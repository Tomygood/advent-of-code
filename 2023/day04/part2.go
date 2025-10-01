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
	table := []int{}
	for i := range len(strings.Split(inputDay, "\n")) {
		table, _ = append(table, 1), i
	}
	for a, game := range strings.Split(inputDay, "\n") {
		var value = 0
		numbers := strings.Split(game, " | ")
		winning := strings.Split(numbers[0], ": ")[1]
		own := numbers[1]
		// fmt.Println(winning)
		// fmt.Println(own)
		for k := 0; k < len(own)-1; k += 3 {
			for i := 0; i < len(winning)-1; i += 3 {
				// fmt.Println(string(own[k])+string(own[k+1]), "r", string(winning[i])+string(winning[i+1]), "r")
				if string(own[k])+string(own[k+1]) == string(winning[i])+string(winning[i+1]) {
					value += 1
					table[a+value] += table[a]
				}
			}
		}
		total += table[a]
	}
	fmt.Println(total)
}
