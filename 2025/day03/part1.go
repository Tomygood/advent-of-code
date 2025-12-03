package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func joliage(bank string) int {
	b := utils.Map(strings.Split(bank, ""), utils.Atoi)

	best := b[0] + b[1]

	for i, a := range b {
		for _, c := range b[i+1:] {
			if utils.ConcatInts(a, c) > best {
				best = utils.ConcatInts(a, c)
			}
		}

	}
	return best

}

func part1() {

	lines := utils.Lines(inputDay)

	var res int
	for _, bank := range lines {
		res += joliage(bank)

	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
