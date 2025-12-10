package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func area(a, b utils.Point) int {
	return (utils.Abs(a.X-b.X) + 1) * (utils.Abs(a.Y-b.Y) + 1)
}

func part1() {

	lines := utils.Lines(inputDay)

	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
	}

	var corners []utils.Point
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		corners = append(corners, utils.Point{X: y, Y: x})
	}

	var marea int
	for i, c := range corners {
		for j := i + 1; j < len(corners); j++ {
			a := area(c, corners[j])
			if a > marea {
				marea = a
			}
		}
	}

	fmt.Println(marea)
	utils.ToClipboard(marea)
}
