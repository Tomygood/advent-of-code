package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y int
}

func dijkstra(start Point, end Point, grid [][]rune) (int, map[Point]Point) {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)
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
	return scores[end], came_from
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
			// n, _ = strconv.Atoi(cha)
			// mat[i][j] = n
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

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)

	star, end := find(grid, 'S'), find(grid, 'E')
	grid[star.x][star.y] = '.'
	grid[end.x][end.y] = '.'

	to_remove := []Point{}
	for x, line := range grid {
		for y, c := range line {
			if x > 0 && y > 0 && x < len(grid)-1 && y < len(grid[0])-1 {
				if c == '#' && ((grid[x+1][y] == '.' && grid[x-1][y] == '.') || (grid[x][y+1] == '.' && grid[x][y-1] == '.')) {
					to_remove = append(to_remove, Point{x, y})
				}
			}
		}
	}

	lgth, _ := dijkstra(star, end, grid)
	res := 0

	for _, w := range to_remove {
		grid[w.x][w.y] = '.'
		n_l, _ := dijkstra(star, end, grid)
		if lgth-n_l >= 100 {
			res++
		}
		grid[w.x][w.y] = '#'
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
