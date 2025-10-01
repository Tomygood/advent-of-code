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

	input := strings.Split(inputDay, " ")
	stones := []int{}
	for _, stone := range input {
		a, _ := strconv.Atoi(stone)
		stones = append(stones, a)
	}

	for j := range 25 {
		for i, stone := range stones {
			if stone == 0 {
				stones[i] = 1
			} else if len(strconv.Itoa(stone))%2 == 0 {
				st := int(math.Pow10(len(strconv.Itoa(stone)) / 2))
				stones[i] = stone / st
				stones = append(stones, stone%st)
			} else {
				stones[i] = stone * 2024
			}
		}
	}

	fmt.Println(len(stones))

	fmt.Printf("Took %s", time.Since(start))
}
