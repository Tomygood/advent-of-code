package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x, y int
}

func makeMatrix(n int) [][]rune {
	matrix := make([][]rune, n)
	for i := range matrix {
		matrix[i] = make([]rune, n)
	}
	return matrix
}

func contains(l []Point, x Point) bool {
	for _, el := range l {
		if el == x {
			return true
		}
	}
	return false
}

func reverse(l []Point) []Point {
	res := []Point{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")

	tra := strings.Split(lines[0], "\r\n")
	grid := makeMatrix(len(tra) * 2)
	robot := Point{0, 0}

	for i, line := range tra {
		for j, cha := range line {
			if cha == 'O' {
				grid[i][j*2], grid[i][j*2+1] = '[', ']'
			} else if cha == '#' {
				grid[i][j*2], grid[i][j*2+1] = '#', '#'
			} else {
				grid[i][j*2], grid[i][j*2+1] = cha, '.'
				if cha == '@' {
					robot.x, robot.y = i, j*2
					grid[i][j*2] = '.'
				}
			}
		}
	}

	moves := ""
	for _, move := range lines[1] {
		if move != '\n' && move != '\r' {
			moves += string(move)
		}
	}

	for _, move := range moves {
		dx, dy := 0, 0
		if move == '^' {
			dx = -1
		} else if move == 'v' {
			dx = 1
		} else if move == '<' {
			dy = -1
		} else {
			dy = 1
		}

		if grid[robot.x+dx][robot.y+dy] == '.' {
			robot.x, robot.y = robot.x+dx, robot.y+dy
		} else if grid[robot.x+dx][robot.y+dy] == '#' {
			continue
		} else {

			to_move := []Point{}
			if grid[robot.x+dx][robot.y+dy] == '[' {
				to_move = append(to_move, Point{robot.x + dx, robot.y + dy})
			} else {
				to_move = append(to_move, Point{robot.x + dx, robot.y + dy - 1})
			}

			push := true

			if dx != 0 {
				finished := false
				for !finished {
					finished = true

					for _, box := range to_move {
						if grid[box.x+dx][box.y] == '[' {
							if !contains(to_move, Point{box.x + dx, box.y}) {
								to_move = append(to_move, Point{box.x + dx, box.y})
								finished = false
							}
						}
						if grid[box.x+dx][box.y-1] == '[' {
							if !contains(to_move, Point{box.x + dx, box.y - 1}) {
								to_move = append(to_move, Point{box.x + dx, box.y - 1})
								finished = false
							}
						}
						if grid[box.x+dx][box.y+1] == '[' {
							if !contains(to_move, Point{box.x + dx, box.y + 1}) {
								to_move = append(to_move, Point{box.x + dx, box.y + 1})
								finished = false
							}
						}
						if grid[box.x+dx][box.y] == '#' || grid[box.x+dx][box.y+1] == '#' {
							push = false
							break
						}
					}
				}
				if push {
					for _, mv := range reverse(to_move) {
						grid[mv.x][mv.y], grid[mv.x][mv.y+1] = '.', '.'
						grid[mv.x+dx][mv.y], grid[mv.x+dx][mv.y+1] = '[', ']'
					}
					robot.x = robot.x + dx
				}

			} else {
				cur := Point{robot.x, robot.y + dy}

				for grid[cur.x][cur.y] != '.' {
					cur.y += dy

					if grid[cur.x][cur.y] == '[' && !contains(to_move, cur) {
						to_move = append(to_move, cur)
					}

					if grid[cur.x][cur.y] == '#' {
						push = false
						break
					}
				}
				if push {
					grid[robot.x][robot.y+dy] = '.'
					for _, mv := range to_move {
						grid[mv.x][mv.y+dy], grid[mv.x][mv.y+dy+1] = '[', ']'
					}
					robot.y = robot.y + dy
				}
			}
		}
	}

	res := 0
	for i, line := range grid {
		for j, c := range line {
			if c == '[' {
				res += 100*i + j
			}
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
