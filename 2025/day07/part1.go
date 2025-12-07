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
	tachyons := make(map[int]struct{}, 1)
	tachyons[start.Y] = struct{}{}
	n := len(grid[0])

	for i := range len(grid) - 1 {
		var ntach = make(map[int]struct{}, min(n, len(tachyons)*2))
		for tach := range tachyons {
			if grid[i][tach] == '^' {
				ntach[tach-1] = struct{}{}
				ntach[tach+1] = struct{}{}
				res++
			} else {
				ntach[tach] = struct{}{}
			}
		}
		tachyons = ntach
	}
	fmt.Println(res)
	utils.ToClipboard(res)
}
