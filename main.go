package main

import (
	"advent-of-code/utils"
	_ "embed"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Part 1: ")
	start := time.Now()
	a := []int{1, 2, 3, 4}
	fmt.Println(utils.Reverse(a))
	fmt.Printf("Took %s\n\n", time.Since(start))

	fmt.Println("Part 2: ")
	start = time.Now()

	fmt.Printf("Took %s\n", time.Since(start))
}
