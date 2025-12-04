package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	lines := utils.ParseMatrix(inputDay)
	var tot int
	res := 1

	for res != 0 {
		res = 0
		lifts := [][2]int{}

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
					lifts = append(lifts, [2]int{i, j})
					res++
				}
			}
		}
		for _, lift := range lifts {
			lines[lift[0]][lift[1]] = 'x'
			tot++
		}
	}

	fmt.Println(tot)
	utils.ToClipboard(tot)
}
