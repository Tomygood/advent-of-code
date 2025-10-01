package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

type Beam struct {
	x, y, dx, dy int
}

func (b Beam) Move(grid [][]rune) []Beam {
	nx, ny := b.x+b.dx, b.y+b.dy

	if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid) {
		return []Beam{}
	}

	if grid[ny][nx] == '.' || (grid[ny][nx] == '-' && b.dx != 0) || (grid[ny][nx] == '|' && b.dy != 0) {
		return []Beam{{nx, ny, b.dx, b.dy}}
	}
	if grid[ny][nx] == '/' {
		return []Beam{{nx, ny, -b.dy, -b.dx}}
	}
	if grid[ny][nx] == '\\' {
		return []Beam{{nx, ny, b.dy, b.dx}}
	}
	if grid[ny][nx] == '-' {
		return []Beam{{nx, ny, 1, 0}, {nx, ny, -1, 0}}
	}
	if grid[ny][nx] == '|' {
		return []Beam{{nx, ny, 0, 1}, {nx, ny, 0, -1}}
	}
	return []Beam{}
}

func makeMatrix(n int) [][]rune {
	matrix := make([][]rune, n)

	for i := range matrix {
		matrix[i] = make([]rune, n)
	}

	return matrix
}

func printMatrix(grid [][]rune, bs []Beam) {
	for _, b := range bs {
		grid[b.y][b.x] = '#'
	}

	for i := range len(grid) {
		for j := range len(grid) {
			print(string(grid[i][j]) + "  ")
		}
		print("\n")
	}
}

func contains(s []Beam, b Beam) bool {
	for _, a := range s {
		if a == b {
			return true
		}
	}
	return false
}

func step(beams []Beam, grid [][]rune, af *[]Beam) []Beam {
	var res = []Beam{}
	for _, b := range beams {
		adv := b.Move(grid)
		for _, nb := range adv {
			if !contains(*af, nb) {
				res = append(res, nb)
				*af = append(*af, nb)
			}
		}
	}
	return res
}

func printBeam(b Beam) {
	fmt.Printf("(%d, %d), (%d, %d)\n", b.x, b.y, b.dx, b.dy)
}

func printBeams(bs []Beam) {
	for _, b := range bs {
		printBeam(b)
	}
}

func copy_matrix(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	return duplicate
}

func main() {
	var grid = strings.Split(inputDay, "\n")
	mat := makeMatrix(len(grid))
	//print(mat)
	for i, line := range grid {
		for j, cha := range strings.TrimSuffix(line, "\n") {
			if j != len(grid) {
				mat[i][j] = cha
			}
		}
	}

	var already_found = []Beam{}
	var beams = []Beam{{-1, 0, 1, 0}}
	energized := makeMatrix(len(grid))

	for i, line := range grid {
		for j := range line {
			if j != len(grid) {
				energized[i][j] = '.'
			}
		}
	}

	for len(beams) != 0 {
		beams = step(beams, mat, &already_found)
		// printBeams(beams)
		// print("\n")
		// if i%2 == 0 {
		// 	printMatrix(copy_matrix(mat), beams)
		// 	print("\n")
		// }
		for _, b := range beams {
			energized[b.y][b.x] = '#'
		}
	}
	var res = 0
	for i, line := range energized {
		for j := range line {
			if energized[i][j] == '#' {
				res += 1
			}
		}
	}
	//printBeams(beams)
	// printMatrix(energized, []Beam{})
	print(res)
}
