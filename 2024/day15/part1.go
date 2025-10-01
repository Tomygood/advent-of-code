package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y int
}

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

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")
	grid := parseMatrix(lines[0])

	moves := ""
	for _, move := range lines[1] {
		if move != '\n' && move != '\r' {
			moves += string(move)
		}
	}

	robot := Point{0, 0}
	for i, line := range grid {
		for j, c := range line {
			if c == '@' {
				robot.x, robot.y = i, j
				grid[i][j] = '.'
			}
		}
	}

	for _, move := range moves {
		dx, dy := 0, 0
		if move == '^' {
			dx = -1
		} else if move == 'v' {
			dx = 1
		} else if move == '<' {
			dy = -1
		} else {
			dy = 1
		}

		if grid[robot.x+dx][robot.y+dy] == '.' {
			robot.x, robot.y = robot.x+dx, robot.y+dy
		} else if grid[robot.x+dx][robot.y+dy] == '#' {
			continue
		} else {
			cur := Point{robot.x + dx, robot.y + dy}
			push := true
			for grid[cur.x][cur.y] != '.' {
				cur.x += dx
				cur.y += dy
				if grid[cur.x][cur.y] == '#' {
					push = false
					break

				}
			}
			if push {
				grid[robot.x+dx][robot.y+dy] = '.'
				grid[cur.x][cur.y] = 'O'

				robot.x, robot.y = robot.x+dx, robot.y+dy
			}
		}
	}
	res := 0
	for i, line := range grid {
		for j, c := range line {
			if c == 'O' {
				res += 100*i + j
			}
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
