package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	a := strings.Split(inputDay, "\r\n\r\n")
	var res int

	var ranges = [][2]int{}
	for _, ingf := range utils.Lines(a[0]) {
		var a, b int
		fmt.Sscanf(ingf, "%d-%d", &a, &b)
		ranges = append(ranges, [2]int{a, b})
	}

	for _, x := range utils.Lines(a[1]) {

		ing := utils.Atoi(x)
		for _, ingf := range ranges {
			if ing >= ingf[0] && ing <= ingf[1] {
				res++
				break
			}
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
