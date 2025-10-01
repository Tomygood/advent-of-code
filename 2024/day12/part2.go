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

type Fence struct {
	x, y, dx, dy int
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

func area(cur Point, seen map[Point]bool, crop rune, grid [][]rune, perim map[Fence]bool, visited map[Point]bool) int {
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
		perim[Fence{cur.x - 1, cur.y, 1, 0}] = true
	}
	if cur.y > 0 {
		if grid[cur.y-1][cur.x] == crop {
			if !seen[Point{cur.x, cur.y - 1}] {
				res += area(Point{cur.x, cur.y - 1}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.y == 0 || grid[cur.y-1][cur.x] != crop {
		perim[Fence{cur.x, cur.y - 1, 0, 1}] = true
	}
	if cur.x < len(grid)-1 {
		if grid[cur.y][cur.x+1] == crop {
			if !seen[Point{cur.x + 1, cur.y}] {
				res += area(Point{cur.x + 1, cur.y}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.x == len(grid)-1 || grid[cur.y][cur.x+1] != crop {
		perim[Fence{cur.x + 1, cur.y, -1, 0}] = true
	}
	if cur.y < len(grid)-1 {
		if grid[cur.y+1][cur.x] == crop {
			if !seen[Point{cur.x, cur.y + 1}] {
				res += area(Point{cur.x, cur.y + 1}, seen, crop, grid, perim, visited)
			}
		}
	}
	if cur.y == len(grid)-1 || grid[cur.y+1][cur.x] != crop {
		perim[Fence{cur.x, cur.y + 1, 0, -1}] = true
	}

	return res + 1
}

func across(s Fence, dx, dy int, perim map[Fence]bool, vis map[Fence]bool, started bool) int {
	if started {
		vis[s] = true
	}
	if perim[Fence{s.x + dx, s.y + dy, s.dx, s.dy}] && !vis[Fence{s.x + dx, s.y + dy, s.dx, s.dy}] {
		return across(Fence{s.x + dx, s.y + dy, s.dx, s.dy}, dx, dy, perim, vis, started)
	}
	if perim[Fence{s.x, s.y, dx, dy}] && !vis[Fence{s.x, s.y, dx, dy}] {
		return across(Fence{s.x, s.y, dx, dy}, -s.dx, -s.dy, perim, vis, true) + 1
	}
	if dx != 0 {
		if perim[Fence{s.x + dx, s.y + s.dy, -dx, -dy}] && !vis[Fence{s.x + dx, s.y + s.dy, -dx, -dy}] {
			return across(Fence{s.x + dx, s.y + s.dy, -dx, -dy}, s.dx, s.dy, perim, vis, true) + 1
		}
	} else {
		if perim[Fence{s.x + s.dx, s.y + dy, 0, -dy}] && !vis[Fence{s.x + s.dx, s.y + dy, 0, -dy}] {
			return across(Fence{s.x + s.dx, s.y + dy, 0, -dy}, s.dx, 0, perim, vis, true) + 1
		}
	}
	return 0
}

func num(cur Point, grid [][]rune, visited map[Point]bool) (int, int) {
	perim := map[Fence]bool{}
	a := area(cur, map[Point]bool{}, grid[cur.y][cur.x], grid, perim, visited)
	p := 0
	vis := map[Fence]bool{}

	for k, _ := range perim {
		if !vis[k] {
			p += across(k, k.dy, k.dx, perim, vis, false)
		}
	}
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
