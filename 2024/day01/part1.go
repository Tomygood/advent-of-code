package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
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
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	res := 0
	for i := range left {
		res += int(math.Abs(float64(right[i] - left[i])))
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
