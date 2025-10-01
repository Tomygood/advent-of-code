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

func is_in_rules(l [][2]int, el [2]int) bool {
	for _, x := range l {
		if x == el {
			return true
		}
	}
	return false
}

func is_valid(pages []int, rules [][2]int) bool {
	for i := 0; i < len(pages); i++ {
		for j := i; j < len(pages); j++ {
			if is_in_rules(rules, [2]int{pages[j], pages[i]}) {
				return false
			}
		}
	}
	return true
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")
	rules_string := strings.Split(lines[0], "\r\n")
	updates := strings.Split(lines[1], "\r\n")

	rules := [][2]int{}
	for _, rule := range rules_string {
		b, a := 0, 0
		fmt.Sscanf(rule, "%d|%d", &b, &a)
		rules = append(rules, [2]int{b, a})
	}

	res := 0
	for _, update := range updates {
		pages_strings := strings.Split(update, ",")
		pages := []int{}
		for _, el := range pages_strings {
			a, err := strconv.Atoi(el)
			if err == nil {
				pages = append(pages, a)
			}
		}
		if is_valid(pages, rules) {
			res += pages[len(pages)/2]
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
