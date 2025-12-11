package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func find_paths(start, end string, conn map[string][]string) int {

	paths := map[string]int{end: 1}

	for paths[start] == 0 {

		for k, v := range conn {

			if paths[k] != 0 {
				continue
			}

			flag := true
			var tot int
			for _, w := range v {
				if paths[w] == 0 {
					flag = false
					break
				}
				tot += paths[w]
			}
			if !flag {
				continue
			}
			paths[k] = tot
		}
	}
	return paths[start]
}

func part1() {

	lines := utils.Lines(inputDay)

	conn := map[string][]string{}

	for _, line := range lines {
		a := strings.Split(line, ":")
		for _, b := range strings.Split(a[1][1:], " ") {
			conn[a[0]] = append(conn[a[0]], b)
		}
	}

	res := find_paths("you", "out", conn)

	fmt.Println(res)
	utils.ToClipboard(res)
}
