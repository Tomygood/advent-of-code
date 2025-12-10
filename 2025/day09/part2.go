package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func rg(a, b utils.Point, corners []utils.Point) bool {
	tl := utils.Point{X: min(a.X, b.X) + 1, Y: min(a.Y, b.Y) + 1}
	br := utils.Point{X: max(a.X, b.X) - 1, Y: max(a.Y, b.Y) - 1}
	for i, c1 := range corners {
		c2 := corners[(i+1)%len(corners)]
		minX := min(c1.X, c2.X)
		maxX := max(c1.X, c2.X)
		minY := min(c1.Y, c2.Y)
		maxY := max(c1.Y, c2.Y)
		if maxY >= tl.Y &&
			minY <= br.Y &&
			maxX >= tl.X &&
			minX <= br.X {
			return true
		}
	}
	return false
}

func part2() {

	lines := utils.Lines(inputDay)

	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
	}

	corners := []utils.Point{}
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		corners = append(corners, utils.Point{X: y, Y: x})
	}
	marea := 0
	for i, c := range corners {
		for j := i + 1; j < len(corners); j++ {
			if !rg(c, corners[j], corners) {
				a := area(c, corners[j])
				if a > marea {
					marea = a
				}
			}

		}
	}

	fmt.Println(marea)
	utils.ToClipboard(marea)
}
