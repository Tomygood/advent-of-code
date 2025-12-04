package utils

import (
	"log"
	"math"
	"strings"
)

// CloneMatrix returns a fresh copy of the given matrix
func CloneMatrix[L ~[][]E, E any](matrix L) [][]E {
	duplicate := make([][]E, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]E, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

// PrintMatrix outputs the given matrix
func PrintMatrix(matrix [][]rune) {
	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			print(string(matrix[i][j]))
			print("  ")
		}
		print("\n")
	}
}

// MakeMatrix[type] returns a brand new n × m matrix with type elements (type HAS to be specified)
func MakeMatrix[E any](n, m int) [][]E {
	matrix := make([][]E, n)
	for i := range matrix {
		matrix[i] = make([]E, m)
	}
	return matrix
}

// ParseIntMatrix parses a generic string into a matrix of integers
func ParseIntMatrix(s string) [][]int {
	var grid = strings.Split(s, "\r\n")
	mat := MakeMatrix[int](len(grid), len(grid[0]))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = Atoi(string(cha))
		}
	}
	return mat
}

// ParseMatrix parses a generic string into a matrix of runes
func ParseMatrix(s string) [][]rune {
	var grid = strings.Split(s, "\r\n")
	mat := MakeMatrix[rune](len(grid), len(grid[0]))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = cha
		}
	}
	return mat
}

func PadMatrix[L ~[][]E, E any](matrix L, val E) L {
	n, m := len(matrix), len(matrix[0])
	res := MakeMatrix[E](n+2, m+2)

	res[0] = make([]E, m+2)
	res[n] = make([]E, m+2)

	for i := range m + 2 {
		res[0][i], res[n+1][i] = val, val
	}

	for i, line := range res[1 : n+1] {
		line[0], line[m+1] = val, val
		for j := range line[1 : m+1] {
			res[i+1][j+1] = matrix[i][j]
		}
	}

	return res
}

// Find returns a Point structure containing the first found occurence of goal in matrix
func Find[L ~[][]E, E comparable](matrix L, goal E) Point {
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

// FindPointD returns a PointD structure containing the first found occurence of goal in matrix
func FindPointD[L ~[][]E, E comparable](matrix L, goal E) PointD {
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

// Homemade slow-aah DFS do NOT USE, dijkstra just below (here for legacy reasons)
func DFS(pos Point, grid [][]rune, seen map[Point]bool, scores map[Point]int, score int, end Point) int {
	if pos == end {
		return score
	}
	min := math.MaxInt

	for _, d := range Deltas {
		p := Point{pos.X + d.X, pos.Y + d.Y}
		best := scores[p]
		if grid[pos.X+d.X][pos.Y+d.Y] != '#' && !seen[p] && (best > score+1) {
			seen[p] = true
			scores[p] = score + 1
			w := DFS(p, grid, seen, scores, score+1, end)
			if w < min {
				min = w
			}
			seen[p] = false
		}
	}

	return min
}

// Path recreates the path taken to go from start to end using the came_from map
func Path(start Point, end Point, came_from map[Point]Point) []Point {
	res := []Point{end}
	cur := end
	for cur != start {
		cur = came_from[end]
		res = append(res, cur)
	}
	return Reverse(res)
}

// Insert inserts Point el in dequeue queue based on points scores
func Insert(queue []Point, scores map[Point]int, el Point) {
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

// Dijkstra returns the length of the shortest path from start to end, as well as a way to recreate paths to each points
func Dijkstra[L ~[][]E, E comparable](start Point, end Point, grid L, wall E) (int, map[Point]Point) {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)

	scores[start] = 0

	pqueue := []Point{start}

	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range Deltas {
			ngh := Point{v.X + d.X, v.Y + d.Y}
			if grid[ngh.X][ngh.Y] != wall {
				_, ok := came_from[ngh]
				new_score := scores[v] + 1
				if !ok || scores[ngh] > new_score {
					came_from[ngh] = v
					scores[ngh] = new_score
					Insert(pqueue, scores, ngh)
				}

			}

		}
	}
	return scores[end], came_from
}

// DijkstraCosts performs Dijkstra’s algorithm but with a cost function
func DijkstraCosts[L ~[][]E, E comparable](start Point, end Point, grid L, wall E, cost func(Point, Point) int) (int, map[Point]Point) {
	came_from := make(map[Point]Point)
	scores := make(map[Point]int)

	scores[start] = 0

	pqueue := []Point{start}

	for len(pqueue) > 0 {
		v := pqueue[0]
		pqueue = pqueue[1:]

		for _, d := range Deltas {
			ngh := Point{v.X + d.X, v.Y + d.Y}
			if grid[ngh.X][ngh.Y] != wall {
				_, ok := came_from[ngh]
				new_score := scores[v] + cost(v, ngh)
				if !ok || scores[ngh] > new_score {
					came_from[ngh] = v
					scores[ngh] = new_score
					Insert(pqueue, scores, ngh)
				}

			}

		}
	}
	return scores[end], came_from
}
