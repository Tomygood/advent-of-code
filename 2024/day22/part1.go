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

	res := 0

	for _, monkey := range lines {
		secret := atoi(monkey)
		for range 2000 {
			secret = process(secret)
		}
		res += secret
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
