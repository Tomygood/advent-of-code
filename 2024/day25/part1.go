package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func parseMatrix(s string) [][]rune {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix(max(len(grid), len(grid[0])))
	for i, line := range grid {
		for j, cha := range line {
			// n, _ = strconv.Atoi(cha)
			// mat[i][j] = n
			mat[i][j] = cha
		}
	}
	return mat
}

func makeMatrix(n int) [][]rune {
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, n)
	}
	return matrix
}

func scoreKey(grid [][]rune) []int {
	res := []int{}
	for y := range len(grid[0]) {
		curx := len(grid) - 1
		for grid[curx][y] == '#' {
			curx--
		}
		res = append(res, len(grid)-2-curx)
	}
	return res
}

func scoreLock(grid [][]rune) []int {
	res := []int{}
	for y := range len(grid[0]) {
		curx := 0
		for grid[curx][y] == '#' {
			curx++
		}
		res = append(res, curx)
	}
	return res
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")
	locksScores := [][]int{}
	keysScores := [][]int{}

	WIDTH := 5
	HEIGHT := 7

	for _, scheme := range lines {
		if scheme[0] == '#' {
			locksScores = append(locksScores, scoreLock(parseMatrix(scheme))[:WIDTH])
		} else {
			keysScores = append(keysScores, scoreKey(parseMatrix(scheme))[:WIDTH])

		}
	}

	res := 0

	for _, lock := range locksScores {
		for _, key := range keysScores {
			valid := true
			for col := range WIDTH {
				if lock[col]+key[col] >= HEIGHT {
					valid = false
				}
			}
			if valid {
				res++
			}
		}
	}

	fmt.Println(res)
	fmt.Printf("Took %s", time.Since(start))
}
