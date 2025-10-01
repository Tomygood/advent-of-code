package utils

import (
	"log"
	"math"
	"strings"
)

func copyMatrix(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func printMatrix(matrix [][]rune) {
	for i := range len(matrix) {
		for j := range len(matrix) {
			print(string(matrix[i][j]))
			print("  ")
		}
		print("\n")
	}
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
				return Point{i, j, 0, 0}
			}
		}
	}
	log.Fatal("Value not found")
	return Point{}
}

// homemade slow-aah DFS do NOT USE, dijkstra just below (here for legacy reasons)
func dfs(pos Point, grid [][]rune, seen map[Point]bool, scores map[Point]int, score int, end Point) int {
	if pos == end {
		return score
	}
	deltas := []Point{Point{1, 0}, Point{-1, 0}, Point{0, 1}, Point{0, -1}}
	min := math.MaxInt

	for _, d := range deltas {
		p := Point{pos.x + d.x, pos.y + d.y}
		best := scores[p]
		if grid[pos.x+d.x][pos.y+d.y] != '#' && !seen[p] && (best > score+1) {
			seen[p] = true
			scores[p] = score + 1
			w := dfs(p, grid, seen, scores, score+1, end)
			if w < min {
				min = w
			}
			seen[p] = false
		}
	}

	return min
}

func reverse(l []Point) []Point {
	res := []Point{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

func path(start Point, end Point, came_from map[Point]Point) []Point {
	res := []Point{end}
	cur := end
	for cur != start {
		cur = came_from[end]
		res = append(res, cur)
	}
	return reverse(res)
}

func insert(queue []Point, scores map[Point]int, el Point) {
	if len(queue) == 0 {
		queue = []Point{el}
	} else {
		start, end := 0, len(queue)-1

		for end-start > 1 {
			m := (end - start) / 2
			if scores[queue[m]] < scores[el] {
				start = m
			} else {
				end = m
			}
		}
		queue = append(queue[:start+1], queue[start:]...)
		queue[start] = el
	}
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
					insert(pqueue, scores, ngh)
				}

			}

		}
	}
	return scores[end], came_from
}
