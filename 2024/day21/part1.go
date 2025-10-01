package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

//go:embed direc.txt
var direc string

//go:embed keypad.txt
var keypad string

type Robot struct {
	x, y int
}

type State struct {
	rob  [3]Robot
	goal string
}

func find(matrix [][]rune, goal rune) Robot {
	for i, line := range matrix {
		for j, c := range line {
			if c == goal {
				return Robot{i, j}
			}
		}
	}
	log.Fatal("Value not found")
	return Robot{}
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

func concatInts(a, b int) int {
	res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return res
}

func numeric(s string) int {
	res := 0
	for _, el := range s {
		val, ok := strconv.Atoi(string(el))
		if ok == nil {
			res = concatInts(res, val)
		}
	}
	return res
}

func step(dir rune, robs Robot) Robot {
	if dir == '^' {
		return Robot{robs.x - 1, robs.y}
	}
	if dir == 'v' {
		return Robot{robs.x + 1, robs.y}
	}
	if dir == '<' {
		return Robot{robs.x, robs.y - 1}
	}
	if dir == '>' {
		return Robot{robs.x, robs.y + 1}
	}

	return robs
}

func advance(dir rune, robs *[3]Robot, direc_grid, keypad_grid [][]rune) rune {
	robs[0] = step(dir, robs[0])
	if dir == 'A' {
		robs[1] = step(direc_grid[robs[0].x][robs[0].y], robs[1])

		if robs[1].x < 0 || robs[1].x > 1 || robs[1].y < 0 || robs[1].y > 2 || (robs[1].x == 0 && robs[1].y == 0) {
			return 'x'
		}

		if direc_grid[robs[0].x][robs[0].y] == 'A' {
			robs[2] = step(direc_grid[robs[1].x][robs[1].y], robs[2])

			if robs[2].x < 0 || robs[2].x > 3 || robs[2].y < 0 || robs[2].y > 2 || (robs[2].x == 3 && robs[2].y == 0) {
				return 'x'
			}

			if direc_grid[robs[1].x][robs[1].y] == 'A' {
				return keypad_grid[robs[2].x][robs[2].y]
			}
		}
	}
	return ' '
}

func solve(robs [3]Robot, goal rune, direc_grid, keypad_grid [][]rune, depth int, seen map[[3]Robot]bool) string {
	if depth > 27 {
		return "x"
	}
	res := []string{}
	nrobs := robs
	l := advance('A', &nrobs, direc_grid, keypad_grid)

	if l == goal {
		return "A"

	} else if l == ' ' && !seen[nrobs] {
		seen[nrobs] = true
		res = append(res, "A"+solve(nrobs, goal, direc_grid, keypad_grid, depth+1, seen))
		seen[nrobs] = false
	}

	if robs[0].x == 1 && robs[0].y != 0 {
		advance('^', &robs, direc_grid, keypad_grid)
		if !seen[robs] {
			seen[robs] = true
			res = append(res, "^"+solve(robs, goal, direc_grid, keypad_grid, depth+1, seen))
			seen[robs] = false
		}
		advance('v', &robs, direc_grid, keypad_grid)
	}
	if robs[0].x == 0 {
		advance('v', &robs, direc_grid, keypad_grid)

		if !seen[robs] {
			seen[robs] = true
			res = append(res, "v"+solve(robs, goal, direc_grid, keypad_grid, depth+1, seen))
			seen[robs] = false
		}
		advance('^', &robs, direc_grid, keypad_grid)
	}
	if robs[0].y > 0 && (robs[0].x != 0 || robs[0].y != 1) {

		advance('<', &robs, direc_grid, keypad_grid)

		if !seen[robs] {
			seen[robs] = true
			res = append(res, "<"+solve(robs, goal, direc_grid, keypad_grid, depth+1, seen))
			seen[robs] = false
		}
		advance('>', &robs, direc_grid, keypad_grid)
	}
	if robs[0].y < 2 {
		advance('>', &robs, direc_grid, keypad_grid)
		if !seen[robs] {
			seen[robs] = true
			res = append(res, ">"+solve(robs, goal, direc_grid, keypad_grid, depth+1, seen))
			seen[robs] = false
		}
		advance('<', &robs, direc_grid, keypad_grid)
	}
	if len(res) == 0 {
		return "x"
	}
	best := "x"
	for _, sol := range res {
		if sol[len(sol)-1] != 'x' && (len(sol) < len(best) || best == "x") {
			best = sol
		}
	}
	return best
}

func try(moves string, robs *[3]Robot, direc_grid, keypad_grid [][]rune) string {
	res := ""
	for _, move := range moves {
		a := advance(move, robs, direc_grid, keypad_grid)
		if a != ' ' && a != 'x' {
			res += string(a)
		}
	}
	return res
}

func main() {
	start := time.Now()

	codes := strings.Split(inputDay, "\r\n")
	direc_grid := parseMatrix(direc)
	keypad_grid := parseMatrix(keypad)

	direc_a := find(direc_grid, 'A')
	keypad_a := find(keypad_grid, 'A')

	ans := 0
	for _, code := range codes {
		robots := [3]Robot{direc_a, direc_a, keypad_a}
		res := ""
		for _, digit := range code {
			seen := map[[3]Robot]bool{}
			seen[robots] = true
			a := solve(robots, digit, direc_grid, keypad_grid, 0, seen)
			res += a
			try(a, &robots, direc_grid, keypad_grid)
		}

		try(res, &[3]Robot{direc_a, direc_a, keypad_a}, direc_grid, keypad_grid)
		ans += len(res) * numeric(code)
	}
	fmt.Println(ans)

	fmt.Printf("Took %s", time.Since(start))
}
