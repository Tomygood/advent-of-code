package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func makeMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func parseMatrix(s string) [][]int {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix(len(grid))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = int(cha)
		}
	}
	return mat
}

func trails(grid [][]int, y, x int, seen map[[2]int]bool) int {
	if !seen[[2]int{x, y}] && grid[y][x] == int('9') {
		seen[[2]int{x, y}] = true
		return 1
	}
	res := 0
	if x > 0 && grid[y][x-1] == grid[y][x]+1 {
		res += trails(grid, y, x-1, seen)
	}
	if x < len(grid)-1 && grid[y][x+1] == grid[y][x]+1 {
		res += trails(grid, y, x+1, seen)
	}
	if y < len(grid)-1 && grid[y+1][x] == grid[y][x]+1 {
		res += trails(grid, y+1, x, seen)
	}
	if y > 0 && grid[y-1][x] == grid[y][x]+1 {
		res += trails(grid, y-1, x, seen)
	}
	return res
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)
	res := 0

	for i, line := range grid {
		for j, c := range line {
			if c == '0' {
				seen := map[[2]int]bool{}
				res += trails(grid, i, j, seen)
			}
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
