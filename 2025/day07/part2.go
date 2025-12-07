package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	grid := utils.ParseMatrix(inputDay)

	start := utils.Find(grid, 'S')

	n := len(grid[0])
	tachyons := make(map[int]int, 1)
	tachyons[start.Y] = 1

	for i := range len(grid) - 1 {

		ntech := make(map[int]int, min(n, len(tachyons)*2))
		for tach, w := range tachyons {
			if grid[i][tach] == '^' {
				ntech[tach-1] += w
				ntech[tach+1] += w
			} else {
				ntech[tach] += w
			}
		}
		tachyons = ntech
	}

	var res int
	for _, tech := range tachyons {
		res += tech
	}
	fmt.Println(res)
	utils.ToClipboard(res)
}
