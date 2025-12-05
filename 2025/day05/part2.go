package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func part2() {

	a := strings.Split(inputDay, "\r\n\r\n")

	f_windows := [][2]int{}
	for _, ing := range utils.Lines(a[0]) {
		var a, b int
		fmt.Sscanf(ing, "%d-%d", &a, &b)
		f_windows = append(f_windows, [2]int{a, b})
	}

	windows := slices.Clone(f_windows)

	var n int

	for n != len(windows) {
		n = len(windows)
		for j, ing := range windows {
			a := ing[0]
			b := ing[1]
			flag := false

			for i, wind := range windows {

				if a < wind[0] && b >= wind[0] {
					windows[i][0] = a
					flag = true
				}
				if b > wind[1] && a <= wind[1] {
					windows[i][1] = b
					flag = true
				}
				if flag {
					if j < n-1 {
						windows = append(windows[:j], windows[j+1:]...)
					} else {
						windows = windows[:n-1]
					}
					break
				}
			}
		}
		windows = utils.NoDupes(windows)
	}

	var res int

	for _, win := range windows {
		res += win[1] - win[0] + 1
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
