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

	circuits := [][]utils.Point3D{}
	for _, box := range boxes {
		circuits = append(circuits, []utils.Point3D{box})
	}

	taken := map[[2]utils.Point3D]bool{}

	var ma, mb utils.Point3D
	var la, lb utils.Point3D

	for len(circuits) != 1 {
		m := math.MaxInt

		for i, a := range boxes {
			for _, b := range boxes[i+1:] {
				if taken[[2]utils.Point3D{a, b}] {
					continue
				}
				d := distance(a, b)
				if d < m {
					m = d
					ma, mb = a, b
				}
			}
		}

		taken[[2]utils.Point3D{ma, mb}] = true

		var c1, c2 int
		for c, circuit := range circuits {
			if slices.Contains(circuit, ma) {
				c1 = c
			}
			if slices.Contains(circuit, mb) {
				c2 = c
			}
		}

		if c1 != c2 {
			circuits[c1] = slices.Concat(circuits[c1], circuits[c2])
			la, lb = ma, mb

			if c2 != len(circuits) {
				circuits = append(circuits[:c2], circuits[c2+1:]...)
			} else {
				circuits = circuits[:c2]
			}
		}
	}

	res := la.X * lb.X

	fmt.Println(res)
	utils.ToClipboard(res)
}
