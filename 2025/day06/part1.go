package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func part1() {

	lines := utils.Lines(inputDay)

	ops := [][]int{{}}

	for i, line := range lines {
		a := strings.Split(line, " ")
		var n int
		for _, d := range a {
			if utils.Atoi(d) != 0 {
				ops[n] = append(ops[n], utils.Atoi(d))
				n++
				if i == 0 {
					ops = append(ops, []int{})
				}
			}
		}
	}

	opera := strings.Split(lines[len(lines)-1], " ")
	var n, res, r int

	for _, op := range opera {
		if len(op) != 0 {
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
					r *= v
				}
				res += r
			}
			n++
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
