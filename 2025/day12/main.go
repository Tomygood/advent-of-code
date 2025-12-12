package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var inputDay string

func main() {

	fmt.Println("Part 1: ")
	start := time.Now()
	part1()
	fmt.Printf("Took %s\n\n", time.Since(start))
}
