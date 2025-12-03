// This version doesn’t take 50 seconds computing every possibility in O(n²)…

package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func joliage3(b []int, start int, size int) int {

	if size == 1 {
		return slices.Max(b[start:])
	}

	best := b[start]
	bi := start

	for i, a := range b[start : len(b)-size+1] {
		if a > best {
			best = a
			bi = start + i
		}
	}
	return utils.ConcatInts(best, joliage3(b, bi+1, size-1))
}

func part2_opt() {

	lines := utils.Lines(inputDay)

	var res int
	for _, bank := range lines {
		res += joliage3(utils.Map(strings.Split(bank, ""), utils.Atoi), 0, 12)
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
