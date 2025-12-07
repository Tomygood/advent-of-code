package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	grid := utils.ParseMatrix(inputDay)

	start := utils.Find(grid, 'S')
	var res int
	tachyons := []utils.Point{start}

	for range len(grid) - 1 {
		var ntech []utils.Point
		for _, tach := range tachyons {
			if grid[tach.X+1][tach.Y] == '^' {
				ntech = append(ntech, utils.Point{X: tach.X + 1, Y: tach.Y - 1},
					utils.Point{X: tach.X + 1, Y: tach.Y + 1})
				res++
			} else {
				ntech = append(ntech, utils.Point{X: tach.X + 1, Y: tach.Y})
			}
		}
		tachyons = utils.NoDupes(ntech)
	}
	fmt.Println(res)
	utils.ToClipboard(res)
}
