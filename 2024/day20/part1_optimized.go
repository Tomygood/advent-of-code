// Faster version of part 1 following the same principle as part 2 (instead of running five million Dijkstras)
package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y int
}

func dijkstra(start Point, end Point, grid [][]rune) map[Point]Point {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)
	for x, line := range grid {
		for y := range line {
			scores[Point{x, y}] = math.MaxInt
		}
	}
	scores[start] = 0

	pqueue := []Point{start}

	deltas := []Point{Point{1, 0}, Point{-1, 0}, Point{0, 1}, Point{0, -1}}
	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range deltas {
			ngh := Point{v.x + d.x, v.y + d.y}
			if grid[ngh.x][ngh.y] != '#' {
				_, ok := came_from[ngh]
				new_score := scores[v] + 1
				if !ok || scores[ngh] > new_score {
					came_from[ngh] = v
					scores[ngh] = new_score
					pqueue = append(pqueue, ngh)
				}
			}
		}
	}
	return came_from
}

func reverse(l []Point) []Point {
	res := []Point{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
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

func find(matrix [][]rune, goal rune) Point {
	for i, line := range matrix {
		for j, c := range line {
			if c == goal {
				return Point{i, j}
			}
		}
	}
	log.Fatal("Value not found")
	return Point{}
}

func path(start Point, end Point, came_from map[Point]Point) []Point {
	res := []Point{end}
	cur := end
	for cur != start {
		cur = came_from[cur]
		res = append(res, cur)
	}

	return reverse(res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)

	star, end := find(grid, 'S'), find(grid, 'E')
	grid[star.x][star.y] = '.'
	grid[end.x][end.y] = '.'

	res := 0
	came_from := dijkstra(star, end, grid)
	path_base := path(star, end, came_from)

	res = 0
	for i, p := range path_base {
		for j, p2 := range path_base {
			ans := abs(p.x-p2.x) + abs(p.y-p2.y)
			if ans <= 2 && -(i+ans-1-j) >= 100 {
				res++
			}

		}

	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
