package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	grid := utils.ParseMatrix(inputDay)

	start := utils.Find(grid, 'S')
	tachyons := []utils.Point{start}

	worlds := make(map[utils.Point]int)
	worlds[start] = 1

	for range len(grid) - 1 {

		var ntech []utils.Point
		for _, tach := range tachyons {

			if grid[tach.X+1][tach.Y] == '^' {

				l := utils.Point{X: tach.X + 1, Y: tach.Y - 1}
				r := utils.Point{X: tach.X + 1, Y: tach.Y + 1}
				ntech = append(ntech, l, r)

				worlds[l] += worlds[tach]
				worlds[r] += worlds[tach]

			} else {
				n := utils.Point{X: tach.X + 1, Y: tach.Y}
				ntech = append(ntech, n)
				worlds[n] += worlds[tach]
			}

		}
		tachyons = utils.NoDupes(ntech)
	}

	var res int
	for _, tech := range tachyons {
		res += worlds[utils.Point{X: tech.X, Y: tech.Y}]
	}
	fmt.Println(res)
	utils.ToClipboard(res)
}
