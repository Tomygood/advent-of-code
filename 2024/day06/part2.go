package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func makeMatrix(n int) [][]rune {
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, n)
	}
	return matrix
}

func parseMatrix(s string) [][]rune {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix(len(grid))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = cha
		}
	}
	return mat
}

func step(grid [][]rune, guard *[4]int) bool {
	nx, ny := guard[0]+guard[3], guard[1]+guard[2]
	if grid[nx][ny] == '#' {
		if guard[2] == 0 {
			guard[2] = -guard[3]
			guard[3] = 0
		} else {
			guard[3] = guard[2]
			guard[2] = 0
		}
		return false
	} else {
		guard[0], guard[1] = nx, ny
		return true
	}
}

func isLoop(obst [2]int, guard [4]int, grid [][]rune) bool {
	grid[obst[0]][obst[1]] = '#'
	seen := map[[4]int]bool{}
	for !seen[guard] && guard[0] >= 1 && guard[0] < len(grid)-1 && guard[1] >= 1 && guard[1] < len(grid[0])-1 {
		seen[guard] = true
		step(grid, &guard)
	}
	grid[obst[0]][obst[1]] = '.'
	return seen[guard]
}

func contains(l [][2]int, x [2]int) bool {
	for _, el := range l {
		if el == x {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)
	guard := [4]int{0, 0, 0, -1}
	for i, line := range grid {
		for j, c := range line {
			if c == '^' {
				guard[0] = i
				guard[1] = j
				grid[i][j] = '.'
			}
		}
	}

	guard2 := guard
	moved := false

	// we first figure out where the obstacle can be placed

	num := [][2]int{}
	for guard[0] >= 1 && guard[0] < len(grid)-1 && guard[1] >= 1 && guard[1] < len(grid[0])-1 {
		moved = step(grid, &guard)
		if moved && !contains(num, [2]int{guard[0], guard[1]}) {
			num = append(num, [2]int{guard[0], guard[1]})
		}
	}

	// then we test all positions!

	res := 0
	for _, obst := range num {
		if isLoop(obst, guard2, grid) {
			res++
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
