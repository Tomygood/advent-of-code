package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1_clean() {

	lines := utils.ParseMatrix(inputDay)
	padded := utils.PadMatrix(lines, '.')
	res := 0

	for i, line := range lines {
		for j, c := range line {

			if c != '@' {
				continue
			}

			var rolls int

			for _, d := range utils.DeltasOrtho {
				if padded[i+d.X+1][j+d.Y+1] == '@' {
					rolls++
				}
			}

			if rolls < 4 {
				res++
			}
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
