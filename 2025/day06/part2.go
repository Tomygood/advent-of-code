package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	lines := utils.Lines(inputDay)

	ops := utils.MakeMatrix[string](len(lines[0]), len(lines[0]))

	start_indexes := []int{}

	for i, c := range lines[len(lines)-1] {
		if c == '+' || c == '*' {
			start_indexes = append(start_indexes, i-1)

		}
	}

	for i, line := range lines {
		n := 0
		for j, d := range line {
			if d == '+' || d == '*' {
				break
			}
			if slices.Contains(start_indexes, j) {
				n++
			} else {
				ops[n][i] += string(d)
			}

		}
	}

	nops := utils.MakeMatrix[string](len(ops), len(ops[0]))

	for n, b := range ops {

		for _, l := range b {

			for i := range l {
				cur := l[i]
				nops[n][i] = nops[n][i] + string(cur)
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

				for _, v := range nops[n] {
					v = strings.TrimLeft(v, " ")
					v = strings.TrimRight(v, " ")

					r += utils.Atoi(v)
				}
				res += r
			case "*":
				r = 1

				for _, v := range nops[n] {
					v = strings.TrimLeft(v, " ")
					v = strings.TrimRight(v, " ")

					if utils.Atoi(v) == 0 {
						continue
					}
					r = r * utils.Atoi(v)
				}

				res += r
			}
			n++

		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
