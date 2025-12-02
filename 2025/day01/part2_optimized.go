package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2_optimized() {

	lines := utils.Lines(inputDay)
	var res int
	dial := 50
	for _, line := range lines {

		n := utils.Atoi(line[1:])
		if line[0] == 'R' {
			if dial+n > 99 {
				res += (dial + n) / 100
			}
			dial = utils.Mod(dial+n, 100)
		} else {
			if dial-n <= 0 {
				res += ((n - dial) / 100) + 1
				if dial == 0 {
					res--
				}
			}
			dial = utils.Mod(dial-n, 100)

		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
