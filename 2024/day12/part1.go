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

func area(cur Point, seen map[Point]bool, crop rune, grid [][]rune, perim *int, visited map[Point]bool) int {
	seen[cur] = true
	visited[cur] = true
	res := 0
	if cur.x > 0 {
		if grid[cur.y][cur.x-1] == crop {
			if !seen[Point{cur.x - 1, cur.y}] {
				res += area(Point{cur.x - 1, cur.y}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.x == 0 || grid[cur.y][cur.x-1] != crop {
		*perim++
	}
	if cur.y > 0 {
		if grid[cur.y-1][cur.x] == crop {
			if !seen[Point{cur.x, cur.y - 1}] {
				res += area(Point{cur.x, cur.y - 1}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.y == 0 || grid[cur.y-1][cur.x] != crop {
		*perim++
	}
	if cur.x < len(grid)-1 {
		if grid[cur.y][cur.x+1] == crop {
			if !seen[Point{cur.x + 1, cur.y}] {
				res += area(Point{cur.x + 1, cur.y}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.x == len(grid)-1 || grid[cur.y][cur.x+1] != crop {
		*perim++
	}
	if cur.y < len(grid)-1 {
		if grid[cur.y+1][cur.x] == crop {
			if !seen[Point{cur.x, cur.y + 1}] {
				res += area(Point{cur.x, cur.y + 1}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.y == len(grid)-1 || grid[cur.y+1][cur.x] != crop {
		*perim++
	}

	return res + 1
}

func num(cur Point, grid [][]rune, visited map[Point]bool) (int, int) {
	p := 0
	a := area(cur, map[Point]bool{}, grid[cur.y][cur.x], grid, &p, visited)
	return a, p
}

func main() {
	start := time.Now()

	lines := parseMatrix(inputDay)

	res := 0
	visited := map[Point]bool{}
	for j, line := range lines {
		for i := range line {
			if !visited[Point{j, i}] {
				a, p := num(Point{j, i}, lines, visited)
				res += a * p
			}
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
