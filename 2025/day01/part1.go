package main

import (
	_ "embed"
	"fmt"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	lines := utils.Lines(inputDay)
	res := 0
	dial := 50
	for _, line := range lines {

		n := utils.Atoi(line[1:])
		if line[0] == 'R' {
			dial = utils.Mod(dial+n, 100)
		} else {
			dial = utils.Mod(dial-n, 100)

		}

		if dial == 0 {
			res++
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
