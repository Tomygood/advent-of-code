// this version uses an array keeping track of where a beam is present is instead of a dictionnary

package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1_opt() {

	grid := utils.ParseMatrix(inputDay)

	start := utils.Find(grid, 'S')
	n := len(grid[0])
	tachyons := make([]bool, n)
	tachyons[start.Y] = true

	var res int

	for i := range len(grid) - 1 {
		ntach := make([]bool, n)
		for j, tach := range tachyons {
			if tach {
				if grid[i][j] == '^' {
					ntach[j-1] = true
					ntach[j+1] = true
					res++
				} else {
					ntach[j] = true
				}
			}
		}
		tachyons = ntach
	}
	fmt.Println(res)
	utils.ToClipboard(res)
}
