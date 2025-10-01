package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func makeMatrix(n int) [][]byte {
	matrix := make([][]byte, n)
	for i := range matrix {
		matrix[i] = make([]byte, n)
	}
	return matrix
}

func parseMatrix(s string) [][]byte {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix(len(grid))
	for i, line := range grid {
		for j := range line {
			mat[i][j] = line[j]
		}
	}
	return mat
}

func step(grid [][]byte, guard *[4]int) {
	nx, ny := guard[0]+guard[3], guard[1]+guard[2]
	if grid[nx][ny] == '#' {
		if guard[2] == 0 {
			guard[2] = -guard[3]
			guard[3] = 0
		} else {
			guard[3] = guard[2]
			guard[2] = 0
		}
	} else {
		guard[0], guard[1] = nx, ny
	}
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)
	guard := [4]int{0, 0, 0, -1}
	for i, line := range grid {
		for j, c := range line {
			if c == '^' {
				guard[0], guard[1] = i, j
				grid[i][j] = '.'
			}
		}
	}

	num := makeMatrix(len(grid))
	num[guard[0]][guard[1]] = '.'

	for guard[0] >= 1 && guard[0] < len(grid)-1 && guard[1] >= 1 && guard[1] < len(grid[0])-1 {
		step(grid, &guard)
		num[guard[0]][guard[1]] = '.'
	}
	res := 0
	for _, line := range num {
		for _, c := range line {
			if c == '.' {
				res++
			}
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
