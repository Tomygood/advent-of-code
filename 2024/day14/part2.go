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

func printMatrix(matrix [][]int) {
	for i := range len(matrix) {
		for j := range len(matrix) {
			if matrix[i][j] == 0 {
				print(".  ")
			} else {
				print(strconv.Itoa(matrix[i][j]))
				print("  ")
			}
		}
		print("\n")
	}
}

func makeMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func adj(grid [][]int) int {
	res := 0
	for y, line := range grid {
		for x := range line {
			if grid[y][x] >= 1 && y+1 < len(grid) && x+1 < len(grid) {
				if grid[y][x+1] == grid[y][x] || grid[y+1][x] == grid[y][x] || grid[y+1][x+1] == grid[y][x] {
					res++
				}
			}
		}
	}
	return res
}

func equ(g1 [][]int, g2 [][]int) bool {
	for i := range len(g1) {
		for j := range len(g2) {
			if g1[i][j] != g2[i][j] {
				return false
			}
		}
	}
	return true
}

type Robot struct {
	px, py, vx, vy int
}

func main() {
	start := time.Now()

	width := 101
	height := 103

	lines := strings.Split(inputDay, "\r\n")

	robots := []Robot{}
	og_grid := makeMatrix(max(height, width))

	for _, robot := range lines {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(robot, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, Robot{px, py, vx, vy})
		og_grid[py][px]++
	}

	grid := makeMatrix(max(height, width))

	i := 1
	for !equ(grid, og_grid) {
		grid = makeMatrix(max(height, width))
		for j, robot := range robots {
			robots[j].px = mod((robot.px + robot.vx), width)
			robots[j].py = mod((robot.py + robot.vy), height)

			grid[robots[j].py][robots[j].px]++
		}

		if adj(grid) > 200 {
			printMatrix(grid)
			fmt.Println(i)
		}
		i++

	}

	fmt.Printf("Took %s", time.Since(start))
}
