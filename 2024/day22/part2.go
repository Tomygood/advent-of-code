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

func process(n int) int {
	s1 := ((n * 64) ^ n) % 16777216
	s2 := ((s1 / 32) ^ s1) % 16777216
	s3 := ((s2 * 2048) ^ s2) % 16777216

	return s3
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func main() {
	start := time.Now()
	lines := strings.Split(inputDay, "\r\n")

	sequenceMap := map[[4]int]int{}

	for _, monkey := range lines {
		secret := atoi(monkey)
		seen := map[[4]int]struct{}{}
		sequence := [4]int{}
		price := secret % 10

		for j := range 2000 {
			newSecret := process(secret)
			newPrice := newSecret % 10
			if j >= 3 {
				sequence = [4]int{sequence[1], sequence[2], sequence[3], newPrice - price}

				if _, ok := seen[sequence]; !ok {
					sequenceMap[sequence] += newPrice
					seen[sequence] = struct{}{}
				}
			} else {
				sequence[j+1] = newPrice - price
			}
			secret = newSecret
			price = newPrice
		}
	}

	res := 0
	for _, v := range sequenceMap {
		if v > res {
			res = v
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
