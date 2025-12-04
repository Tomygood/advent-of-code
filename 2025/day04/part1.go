package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	lines := utils.ParseMatrix(inputDay)
	res := 0

	for i, line := range lines {
		for j, c := range line {

			if c != '@' {
				continue
			}

			var rolls int
			if i > 0 {
				// above
				if lines[i-1][j] == '@' {
					rolls++
				}

				if j > 0 {
					if lines[i-1][j-1] == '@' {
						rolls++
					}
				}

				if j < len(lines[0])-1 {
					if lines[i-1][j+1] == '@' {
						rolls++
					}
				}
			}

			if j > 0 {
				if lines[i][j-1] == '@' {
					rolls++
				}
			}

			if j < len(lines[0])-1 {
				if lines[i][j+1] == '@' {
					rolls++
				}
			}

			// below
			if i < len(lines)-1 {
				if lines[i+1][j] == '@' {
					rolls++
				}

				if j > 0 {
					if lines[i+1][j-1] == '@' {
						rolls++
					}
				}

				if j < len(lines[0])-1 {
					if lines[i+1][j+1] == '@' {
						rolls++
					}
				}
			}
			if rolls < 4 {
				fmt.Println(i, j)
				res++
			}
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
