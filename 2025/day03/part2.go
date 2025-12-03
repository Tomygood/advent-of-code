package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func joliage2(bank string, size int) int {
	b := utils.Map(strings.Split(bank, ""), utils.Atoi)

	if size == 1 {
		return slices.Max(b)
	}

	best_i := b[0]
	best := utils.ConcatInts(b[0], joliage2(bank[1:], size-1))

	for i, a := range b[:len(b)-size+1] {
		if a > best_i {
			best_i = a
			best = utils.ConcatInts(a, joliage2(bank[i+1:], size-1))
		}
	}
	return best
}
func part2() {

	lines := utils.Lines(inputDay)

	var res int
	for _, bank := range lines {
		res += joliage2(bank, 12)

	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
