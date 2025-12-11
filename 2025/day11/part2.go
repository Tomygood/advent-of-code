package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func find_paths2(start, end string, conn map[string][]string) int {

	paths := map[string]int{end: 1}

	valid_dac := map[string]int{}
	valid := map[string]int{}

	for paths[start] == 0 {

		for k, v := range conn {

			if paths[k] != 0 {
				continue
			}

			flag := true
			var tot, tot_d, tot_f int
			for _, w := range v {
				if paths[w] == 0 {
					flag = false
					break
				}
				tot += paths[w]
				tot_d += valid_dac[w]
				tot_f += valid[w]
			}

			if !flag {
				continue
			}

			paths[k] = tot
			valid_dac[k] = tot_d
			valid[k] = tot_f
			if k == "dac" {
				valid_dac[k] = tot
			}

			if k == "fft" {
				valid[k] = tot_d
			}

		}
	}
	return valid[start]
}

func part2() {

	lines := utils.Lines(inputDay)

	conn := map[string][]string{}

	for _, line := range lines {
		a := strings.Split(line, ":")
		for _, b := range strings.Split(a[1][1:], " ") {
			conn[a[0]] = append(conn[a[0]], b)
		}
	}

	res := find_paths2("svr", "out", conn)

	fmt.Println(res)
	utils.ToClipboard(res)
}
