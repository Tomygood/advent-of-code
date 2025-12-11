package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func find_paths2(start, end string, conn map[string][]string) int {

	paths := make(map[string]int, len(conn))
	paths[end] = 1

	valid := make(map[string]int, len(conn))

	for paths[start] == 0 {
		for k, v := range conn {
			flag := true
			var tot, tot_f int
			for _, w := range v {
				if paths[w] == 0 {
					flag = false
					break
				}
				tot += paths[w]
				tot_f += valid[w]
			}
			if !flag {
				continue
			}
			paths[k] = tot

			if k == "dac" {
				valid[k] = tot
				continue
			}

			if k == "fft" {
				clear(valid)
			}

			valid[k] = tot_f

			delete(conn, k)
		}
	}
	return valid[start]
}

func part2() {

	lines := utils.Lines(inputDay)

	conn := map[string][]string{}

	for _, line := range lines {
		a := strings.Split(line, ":")
		conn[a[0]] = append(conn[a[0]], strings.Split(a[1][1:], " ")...)
	}

	res := find_paths2("svr", "out", conn)

	fmt.Println(res)
	utils.ToClipboard(res)
}
