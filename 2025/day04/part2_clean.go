package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2_clean() {

	lines := utils.ParseMatrix(inputDay)
	padded := utils.PadMatrix(lines, '.')

	var tot int
	res := 1

	for res != 0 {
		var res int
		var lifted []utils.Point
		for i, line := range lines {
			for j := range line {

				if padded[i+1][j+1] != '@' {
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
					lifted = append(lifted, utils.Point{X: i, Y: j})
				}
			}
		}

		for _, lift := range lifted {
			padded[lift.X+1][lift.Y+1] = '.'
			tot++
		}
	}

	fmt.Println(tot)
	utils.ToClipboard(tot)
}
