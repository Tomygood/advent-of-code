package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	lines := utils.Lines(inputDay)

	ops := utils.MakeMatrix[int](len(lines[0]), len(lines[0]))

	for i, line := range lines {
		a := strings.Split(line, " ")
		n := 0
		for _, d := range a {
			if utils.Atoi(d) != 0 {
				ops[n][i] = utils.Atoi(d)
				n++
			}
		}
	}

	opera := strings.Split(lines[len(lines)-1], " ")
	n := 0
	res := 0

	for _, op := range opera {
		if len(op) != 0 {

			r := 0
			switch op {
			case "+":
				r = 0
				for _, v := range ops[n] {
					r += v
				}
				res += r
			case "*":
				r = 1
				for _, v := range ops[n] {
					if v == 0 {
						break
					}
					r = r * v
				}
				res += r
			}
			n++

		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
