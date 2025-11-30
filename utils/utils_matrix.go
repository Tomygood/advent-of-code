package utils

import (
	"log"
	"math"
	"strconv"
	"strings"
)

// cloneMatrix returns a copy of the given matrix
func cloneMatrix[L ~[][]E, E any](matrix L) [][]E {
	duplicate := make([][]E, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]E, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

// printMatrix outputs the given matrix
func printMatrix(matrix [][]rune) {
	for i := range len(matrix) {
		for j := range len(matrix) {
			print(string(matrix[i][j]))
			print("  ")
		}
		print("\n")
	}
}

// makeMatrix[type] returns a brand new n×n matrix with type elements (type HAS to be specified)
func makeMatrix[E any](n int) [][]E {
	matrix := make([][]E, n)
	for i := range matrix {
		matrix[i] = make([]E, n)
	}
	return matrix
}

// parseIntMatrix parses a generic string into a matrix of integers
func parseIntMatrix(s string) [][]int {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix[int](len(grid))
	for i, line := range grid {
		for j, cha := range line {
			n, _ := strconv.Atoi(string(cha))
			mat[i][j] = n
		}
	}
	return mat
}

// parseMatrix parses a generic string into a matrix of runes
func parseMatrix(s string) [][]rune {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix[rune](len(grid))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = cha
		}
	}
	return mat
}

// find returns a Point structure containing the first found occurence of goal in matrix
func find[L ~[][]E, E comparable](matrix L, goal E) Point {
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

// find returns a PointD structure containing the first found occurence of goal in matrix
func findPointD[L ~[][]E, E comparable](matrix L, goal E) PointD {
	for i, line := range matrix {
		for j, c := range line {
			if c == goal {
				return PointD{i, j, 0, 0}
			}
		}
	}
	log.Fatal("Value not found")
	return PointD{}
}

// homemade slow-aah DFS do NOT USE, dijkstra just below (here for legacy reasons)
func dfs(pos Point, grid [][]rune, seen map[Point]bool, scores map[Point]int, score int, end Point) int {
	if pos == end {
		return score
	}
	deltas := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
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

// path recreates the path taken to go from start to end using the came_from map
func path(start Point, end Point, came_from map[Point]Point) []Point {
	res := []Point{end}
	cur := end
	for cur != start {
		cur = came_from[end]
		res = append(res, cur)
	}
	return reverse(res)
}

// insert inserts Point el in dequeue queue based on points scores
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

// dijkstra returns the length of the shortest path from start to end, as well as a way to recreate paths to each points
func dijkstra[L ~[][]E, E comparable](start Point, end Point, grid L, wall E) (int, map[Point]Point) {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)

	scores[start] = 0

	pqueue := []Point{start}

	deltas := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range deltas {
			ngh := Point{v.x + d.x, v.y + d.y}
			if grid[ngh.x][ngh.y] != wall {
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

// dijkstraCosts performs Dijkstra’s algorithm but with a cost function
func dijkstraCosts[L ~[][]E, E comparable](start Point, end Point, grid L, wall E, cost func(Point, Point) int) (int, map[Point]Point) {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)

	scores[start] = 0

	pqueue := []Point{start}

	deltas := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range deltas {
			ngh := Point{v.x + d.x, v.y + d.y}
			if grid[ngh.x][ngh.y] != wall {
				_, ok := came_from[ngh]
				new_score := scores[v] + cost(v, ngh)
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
