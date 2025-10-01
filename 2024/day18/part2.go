package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
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

			sort.Slice(pqueue, func(i, j int) bool {
				return scores[pqueue[i]] < scores[pqueue[j]]
			})
		}
	}
	return scores[end], came_from
}

func is_on_path(end Point, came_from map[Point]Point, el Point) bool {
	if len(came_from) == 0 {
		return true
	}
	cur := end
	for (cur != Point{1, 1}) {
		cur, _ = came_from[cur]
		if cur == el {
			return true
		}
	}
	return false
}

func makeMatrix(n int) [][]rune {
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, n)
	}
	return matrix
}

func main() {
	start := time.Now()

	const (
		SIZE  = 70
		BYTES = 1024
	)

	grid := makeMatrix(SIZE + 3)
	bpath := map[Point]bool{}
	for x, line := range grid {
		for y := range line {
			bpath[Point{x, y}] = true
			if y == 0 || x == 0 || y == len(grid)-1 || x == len(grid)-1 {
				grid[x][y] = '#'
			} else {
				grid[x][y] = '.'
			}
		}
	}

	lines := strings.Split(inputDay, "\r\n")
	for i, line := range lines {
		if i < BYTES {
			x, y := 0, 0
			fmt.Sscanf(line, "%d,%d", &x, &y)
			grid[y+1][x+1] = '#'
		}
	}
	came_from := map[Point]Point{}
	lgth := 0
	for i, line := range lines {
		if i >= BYTES {
			x, y := 0, 0
			fmt.Sscanf(line, "%d,%d", &x, &y)
			grid[y+1][x+1] = '#'

			if is_on_path(Point{SIZE + 1, SIZE + 1}, came_from, Point{y + 1, x + 1}) {
				lgth, came_from = dijkstra(Point{1, 1}, Point{SIZE + 1, SIZE + 1}, grid)
				if lgth > SIZE*SIZE {
					fmt.Println(strconv.Itoa(x) + "," + strconv.Itoa(y))
					break
				}
			}
		}
	}
	fmt.Printf("Took %s", time.Since(start))
}
