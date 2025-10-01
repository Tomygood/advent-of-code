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

func possibilities(tab []string, tv int) []int {
	res := []int{}
	t, _ := strconv.Atoi(tab[0])
	if len(tab) == 1 {
		return []int{t}
	} else {
		next := possibilities(tab[1:], tv)
		for _, n := range next {
			conc, _ := strconv.Atoi(strconv.Itoa(n) + tab[0])
			if conc <= tv {
				res = append(res, conc)
			}
			if t+n <= tv {
				res = append(res, t+n)
			}
			if t*n <= tv {
				res = append(res, t*n)
			}
		}
	}
	return res
}

func contains(l []int, x int) bool {
	for _, el := range l {
		if el == x {
			return true
		}
	}
	return false
}

func reverse(l []string) []string {
	res := []string{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

func main() {
	start := time.Now()

	equations := strings.Split(inputDay, "\r\n")
	res := 0
	for _, equation := range equations {
		nums := strings.Split(equation, ": ")
		test_value, _ := strconv.Atoi(nums[0])

		if contains(possibilities(reverse(strings.Split(equation, " ")[1:]), test_value), test_value) {
			res += test_value
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
