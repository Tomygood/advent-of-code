// Failed attempt at optimizing more using a sliding window method (runs slower than part2_optimized.go)
package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func joliage4(b []int) int {

	joliage := []int{b[0]}

	nine_index := 0

	for i, v := range b {
		if i == 0 {
			continue
		}
		if i > len(b)-12+len(joliage)-2 {
			if len(joliage) == 12 && v > joliage[11] {
				joliage[11] = v
			}
			break
		}
		flag := false
		for j, f := range joliage[nine_index:] {
			if v > f {
				joliage[j+nine_index] = v
				joliage = joliage[:j+nine_index+1]
				if v == 9 {
					nine_index++
				}
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		if len(joliage) < 12 {
			joliage = append(joliage, v)
		}
	}

	if len(joliage) < 12 {
		joliage = append(joliage, b[len(b)-12+len(joliage):]...)
	}
	return utils.ConcatIntSlice(joliage)
}

func part2_opt2() {

	lines := utils.Lines(inputDay)

	var res int
	for _, bank := range lines {
		res += joliage4(utils.Map(strings.Split(bank, ""), utils.Atoi))
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
