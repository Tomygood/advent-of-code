package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func valid(s string) bool {

	for i := range len(s) {
		if s[i:] == s[:i] {
			return false
		}
	}
	return true
}

func part1() {

	var res int

	s := strings.Split(inputDay, ",")

	for _, a := range s {
		ids := strings.Split(a, "-")

		i := utils.Atoi(ids[0])
		for i <= utils.Atoi(ids[1]) {
			if !valid(strconv.Itoa(i)) {
				res += i
			}
			i++
		}
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
