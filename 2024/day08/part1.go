package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

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

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)
	antennas := map[rune][][2]int{}
	antinode := makeMatrix(len(grid))

	for i, line := range grid {
		for j, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], [2]int{i, j})
			}
		}
	}

	for _, tab := range antennas {
		for i, antenna := range tab {
			for j, antenna2 := range tab {
				if j > i {

					h := antenna2[1] - antenna[1]
					v := antenna2[0] - antenna[0]

					if antenna[1]-h >= 0 && antenna[1]-h < len(antinode) && antenna[0]-v >= 0 && antenna[0]-v < len(antinode) {
						antinode[antenna[0]-v][antenna[1]-h] = '#'
					}
					if antenna2[1]+h >= 0 && antenna2[1]+h < len(antinode) && antenna2[0]+v >= 0 && antenna2[0]+v < len(antinode) {
						antinode[antenna2[0]+v][antenna2[1]+h] = '#'
					}
				}
			}
		}
	}

	res := 0
	for _, line := range antinode {
		for _, char := range line {
			if char == '#' {
				res++
			}
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
