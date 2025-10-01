package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n")
	left, right := []int{}, []int{}
	for _, line := range lines {
		l, r := 0, 0
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}
	m := map[int]int{}
	for _, n := range right {
		_, ok := m[n]
		if ok {
			m[n]++
		} else {
			m[n] = 1
		}
	}

	res := 0
	for _, k := range left {
		res += k * m[k]
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
