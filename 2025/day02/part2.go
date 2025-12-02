package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func valid2(s string) bool {

	for i := range len(s) {

		if i == 0 || (len(s)%i != 0) {
			continue
		}
		flag := false

		for j := range len(s) {

			if (i*(j+2))-1 >= len(s) {
				break
			}
			if j == 0 {
				flag = true
			}

			if s[j*i:i*(j+1)] != s[i*(j+1):i*(j+2)] {
				flag = false
			}
		}

		if flag {
			return false
		}
	}
	return true
}

func part2() {

	var res int

	s := strings.Split(inputDay, ",")

	for _, a := range s {
		ids := strings.Split(a, "-")

		i := utils.Atoi(ids[0])
		for i <= utils.Atoi(ids[1]) {
			if !valid2(strconv.Itoa(i)) {
				res += i
				i += 9
			}
			i++
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
