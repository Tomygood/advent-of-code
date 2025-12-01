package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	lines := utils.Lines(inputDay)
	res := 0
	dial := 50
	for _, line := range lines {

		n := utils.Atoi(line[1:])
		if line[0] == 'R' {
			for range n {
				dial = utils.Mod(dial+1, 100)
				if dial == 0 {
					res++
				}
			}

		} else {
			for range n {
				dial = utils.Mod(dial-1, 100)
				if dial == 0 {
					res++
				}
			}
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
