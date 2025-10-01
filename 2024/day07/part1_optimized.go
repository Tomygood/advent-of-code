// different approach starting from the testing value, idea given by Kévin BUI

package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func is_valid(tv int, tab []string, index int) bool {
	t, _ := strconv.Atoi(tab[index])
	if tv < 0 {
		return false
	}
	if index == 0 {
		return tv == t
	} else {
		return is_valid(tv-t, tab, index-1) || (tv%t == 0 && is_valid(tv/t, tab, index-1))
	}
}

func main() {
	start := time.Now()

	equations := strings.Split(inputDay, "\r\n")
	res := 0
	for _, equation := range equations {
		nums := strings.Split(equation, ": ")
		test_value, _ := strconv.Atoi(nums[0])

		tab := strings.Split(equation, " ")[1:]
		if is_valid(test_value, tab, len(tab)-1) {
			res += test_value
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s\n", time.Since(start))

}
