package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y, dx, dy int
}

type PointD struct {
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

func find(matrix [][]rune, goal rune) Point {
	for i, line := range matrix {
		for j, c := range line {
			if c == goal {
				return Point{i, j, 0, 0}
			}
		}
	}
	log.Fatal("Value not found")
	return Point{}
}

func dijkstra(start Point, end Point, grid [][]rune) int {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)
	deltas := [4]PointD{PointD{1, 0}, PointD{-1, 0}, PointD{0, 1}, PointD{0, -1}}

	for x, line := range grid {
		for y := range line {
			for _, d := range deltas {
				scores[Point{x, y, d.x, d.y}] = math.MaxInt
			}
		}
	}
	scores[start] = 0

	pqueue := []Point{start}
	new_score := 0

	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range deltas {
			ngh := Point{v.x + d.x, v.y + d.y, d.x, d.y}
			if grid[ngh.x][ngh.y] != '#' {
				_, ok := came_from[ngh]
				if d.x != v.dx || d.y != v.dy {
					new_score = scores[v] + 1001
				} else {
					new_score = scores[v] + 1
				}
				if !ok || scores[ngh] > new_score {
					came_from[ngh] = v
					scores[ngh] = new_score
					pqueue = append(pqueue, ngh)
				}

			}

			sort.Slice(pqueue, func(i, j int) bool {
				return scores[pqueue[i]] < scores[pqueue[j]]
			})
		}
	}
	res := math.MaxInt
	for _, d := range deltas {
		a := scores[Point{end.x, end.y, d.x, d.y}]
		if a < res {
			res = a
		}
	}
	return res
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)

	star := find(grid, 'S')
	star.dy = 1
	fmt.Println(dijkstra(star, find(grid, 'E'), grid))

	fmt.Printf("Took %s", time.Since(start))
}
