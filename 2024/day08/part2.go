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
					lh := antenna[1]
					lv := antenna[0]

					for lh >= 0 && lh < len(antinode) && lv >= 0 && lv < len(antinode) {
						antinode[lv][lh] = '#'
						lh -= h
						lv -= v
					}

					rh := antenna2[1]
					rv := antenna2[0]

					for rh >= 0 && rh < len(antinode) && rv >= 0 && rv < len(antinode) {
						antinode[rv][rh] = '#'
						rh += h
						rv += v
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
