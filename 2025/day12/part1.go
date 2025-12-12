package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func ParseMatrix(s string) [][]rune {
	var grid = strings.Split(s, "\r\n")[1:]
	mat := utils.MakeMatrix[rune](len(grid), len(grid[0]))
	for i, line := range grid {
		for j, cha := range line {
			mat[i][j] = cha
		}
	}
	return mat
}

func fill(grid [][]rune, shape [][]rune, i, j int) ([][]rune, bool) {
	ng := utils.CloneMatrix(grid)
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if shape[a][b] != '.' {
				if ng[i+a][j+b] != '.' {
					return ng, false
				}
				ng[i+a][j+b] = shape[a][b]
			}
		}
	}
	return ng, true
}

func rotate(shape [][]rune, rot, flip int) [][]rune {
	nshape := utils.CloneMatrix(shape)

	for range rot {
		turn90(nshape)
	}

	if flip == 1 {
		for j := range 3 {
			nshape[0][j], nshape[2][j] = nshape[2][j], nshape[0][j]
		}
	}

	return nshape
}

func flip_board_v(grid [][]rune) [][]rune {
	ng := utils.CloneMatrix(grid)

	n, m := len(grid), len(grid[0])

	for i := range n / 2 {
		for j := range m {
			ng[i][j], ng[n-1-i][j] = ng[n-1-i][j], ng[i][j]
		}
	}

	return ng
}

func flip_board_h(grid [][]rune) [][]rune {
	ng := utils.CloneMatrix(grid)

	n, m := len(grid), len(grid[0])

	for j := range m / 2 {
		for i := range n {
			ng[i][j], ng[i][m-j-1] = ng[i][m-j-1], ng[i][j]
		}
	}

	return ng
}

func turn90(shape [][]rune) {
	for i := range 3 {
		for j := i; j < 3; j++ {
			shape[i][j], shape[j][i] = shape[j][i], shape[i][j]
		}
	}

	for j := range 3 {
		shape[0][j], shape[2][j] = shape[2][j], shape[0][j]
	}
}

func can_fit(grid [][]rune, to_fit map[int]int, shapes [][][][]rune) bool {
	n, m := len(grid), len(grid[0])
	seen := make(map[string]bool)

	// Check if space is big enough for all presents
	var total int
	for _, v := range to_fit {
		total += v
	}
	if total*8 > n*m {
		return false
	}

	return can_fit_rec(grid, to_fit, shapes, seen, n, m)
}

func can_fit_rec(grid [][]rune, to_fit map[int]int, shapes [][][][]rune, seen map[string]bool, n, m int) bool {

	if len(to_fit) == 0 {
		return true
	}

	// select 1 shape to make fit
	for shape := range to_fit {
		to_fit[shape]--
		if to_fit[shape] == 0 {
			delete(to_fit, shape)
		}

		for i := 0; i+2 < n; i++ {
			for j := 0; j+2 < m; j++ {
				for _, nshape := range shapes[shape] {
					ng, possible := fill(grid, nshape, i, j)

					if !possible {
						continue
					}
					if seen[createKey(ng, to_fit)] {
						continue
					}

					if can_fit_rec(ng, to_fit, shapes, seen, n, m) {
						return true
					}
				}

			}
		}
		to_fit[shape]++
		break
	}
	seen[createKey(grid, to_fit)] = true
	seen[createKey(flip_board_h(grid), to_fit)] = true
	v := flip_board_v(grid)
	seen[createKey(v, to_fit)] = true
	seen[createKey(flip_board_h(v), to_fit)] = true

	return false
}

func createKey(s [][]rune, m map[int]int) string { return fmt.Sprint(s) + fmt.Sprint(m) }

func part1() {

	lines := strings.Split(inputDay, "\r\n\r\n")

	shapes := [][][][]rune{}
	for _, section := range lines[:len(lines)-1] {

		base := ParseMatrix(section)
		rotations := [][][]rune{}
		for rot := range 4 {
			for flip := range 2 {
				nshape := rotate(base, rot, flip)

				genflag := false
				for _, existing_rotation := range rotations {
					flag := true
					for i, line := range existing_rotation {
						if !slices.Equal(line, nshape[i]) {
							flag = false
						}
					}
					if flag {
						genflag = true
					}
				}
				if genflag {
					continue
				}
				rotations = append(rotations, nshape)
			}
		}

		shapes = append(shapes, rotations)
	}

	t := lines[len(lines)-1]
	var res int
	for _, tree := range strings.Split(t, "\r\n") {
		a := strings.Split(tree, ":")
		var n, m int
		fmt.Sscanf(a[0], "%dx%d", &n, &m)

		grid := utils.MakeMatrix[rune](n, m)
		for i := range n {
			for j := range m {
				grid[i][j] = '.'
			}
		}

		then := strings.Split(a[1][1:], " ")
		to_fit := map[int]int{}
		for k, v := range then {
			if v != "0" {
				to_fit[k] = utils.Atoi(v)
			}
		}

		result := can_fit(grid, to_fit, shapes)

		if result {
			res++
		}

	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
