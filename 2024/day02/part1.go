package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func safe(nums []int) bool {
	isz := nums[0] <= nums[1]
	for i := 1; i < len(nums); i++ {
		if math.Abs(float64(nums[i-1]-nums[i])) > 3 || nums[i-1] <= nums[i] != isz || nums[i-1] == nums[i] {
			return false
		}
	}
	return true
}

func str_to_int(l []string) []int {
	t := []int{}
	for _, el := range l {
		v, _ := strconv.Atoi(el)
		t = append(t, v)
	}
	return t
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n")
	safecount := 0
	for _, line := range lines {
		if safe(str_to_int(strings.Split(line, " "))) {
			safecount++
		}
	}
	fmt.Println(safecount)

	fmt.Printf("Took %s", time.Since(start))
}
