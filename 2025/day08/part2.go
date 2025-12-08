package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	lines := utils.Lines(inputDay)

	boxes := []utils.Point3D{}
	for _, b := range lines {
		var x, y, z int
		fmt.Sscanf(b, "%d,%d,%d", &x, &y, &z)
		boxes = append(boxes, utils.Point3D{X: x, Y: y, Z: z})
	}

	circuits := [][]int{}
	circuit_map := map[int]int{}
	for i := range boxes {
		circuits = append(circuits, []int{i})
		circuit_map[i] = i
	}

	taken := map[[2]int]bool{}

	var ma, mb int
	var la, lb int

	cir := len(circuits)
	for cir != 1 {
		m := math.MaxInt

		for i, a := range boxes {
			for j := i + 1; j < len(boxes); j++ {
				b := boxes[j]
				if taken[[2]int{i, j}] {
					continue
				}
				d := distance(a, b)
				if d < m {
					m = d
					ma, mb = i, j
				}
			}
		}

		taken[[2]int{ma, mb}] = true

		c1 := circuit_map[ma]
		c2 := circuit_map[mb]

		if c1 != c2 {
			circuits[c1] = slices.Concat(circuits[c1], circuits[c2])
			la, lb = ma, mb

			for _, c := range circuits[c2] {
				circuit_map[c] = c1
			}

			cir--
		}
	}

	res := boxes[la].X * boxes[lb].X

	fmt.Println(res)
	utils.ToClipboard(res)
}
