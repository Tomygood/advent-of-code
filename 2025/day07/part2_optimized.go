// this version uses an array keeping track of how many beams are present at each point instead of a dictionnary

package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2_opt() {

	grid := utils.ParseMatrix(inputDay)

	start := utils.Find(grid, 'S')

	n := len(grid[0])
	tachyons := make([]int, n)
	tachyons[start.Y] = 1

	for i := range len(grid) - 1 {

		ntech := make([]int, n)
		for tach, w := range tachyons {
			if w != 0 {
				if grid[i][tach] == '^' {
					ntech[tach-1] += w
					ntech[tach+1] += w
				} else {
					ntech[tach] += w
				}
			}
		}
		tachyons = ntech
	}

	var res int
	for _, a := range tachyons {

		res += a
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
