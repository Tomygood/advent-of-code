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

func main() {
	start := time.Now()

	m := map[int]int{}
	for _, stone := range strings.Split(inputDay, " ") {
		a, _ := strconv.Atoi(stone)
		m[a]++
	}
	m2 := map[int]int{}

	for range 75 {
		for k, v := range m {
			if k == 0 {
				m2[1] += v
			} else if len(strconv.Itoa(k))%2 == 0 {
				s := int(math.Pow10(len(strconv.Itoa(k)) / 2))
				m2[k/s] += v
				m2[k%s] += v
			} else {
				m2[k*2024] += v
			}
		}
		m, m2 = m2, map[int]int{}
	}

	res := 0
	for _, v := range m {
		res += v
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
