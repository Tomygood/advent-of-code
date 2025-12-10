// unfinished

package main

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2_routines() {

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

	var marea int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i, c := range corners {

		wg.Add(1)
		go func(valX utils.Point) {
			defer wg.Done()

			var localMax int
			foundCandidate := false

			for j := i + 1; j < len(corners); j++ {
				if !rg(valX, corners[j], corners) {
					a := area(c, corners[j])
					if a > marea {
						marea = a
					}
				}
			}
		}(c)

	}

	fmt.Println(marea)
	utils.ToClipboard(marea)
}
