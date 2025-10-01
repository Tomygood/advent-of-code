// should clean this at some point i don’t think it’s very readable in its current state
// should also optimize this at some point, it’s not very fast

package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed test.txt
var inputDay string

//go:embed direc.txt
var direc string

//go:embed keypad.txt
var keypad string

type Robot struct {
	x, y int
}

type State struct {
	rob  [1]Robot
	goal rune
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
			// n, _ = strconv.Atoi(cha)
			// mat[i][j] = n
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

func advance2(dir rune, robs *[1]Robot, direc_grid [][]rune) rune {
	robs[0] = step(dir, robs[0])
	if robs[0].x < 0 || robs[0].x > 1 || robs[0].y < 0 || robs[0].y > 2 || (robs[0].x == 0 && robs[0].y == 0) {
		return 'x'
	}
	if dir == 'A' {
		return direc_grid[robs[0].x][robs[0].y]
	}
	return ' '
}

func try2(moves string, robs *[1]Robot, direc_grid [][]rune) string {
	res := ""
	for _, move := range moves {
		a := advance2(move, robs, direc_grid)
		if a != ' ' && a != 'x' {
			res += string(a)
		}
	}
	return res
}

func next_a(s string, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == 'A' {
			return i
		}
	}
	return -1
}

func occurences(l []rune) map[rune]int {
	res := map[rune]int{}
	for _, el := range l {
		res[el]++
	}
	return res
}

func sortstr(s string) string {
	res := []rune{}
	for _, el := range s {
		res = append(res, el)
	}
	// occ := occurences(res)
	sort.Slice(res, func(i, j int) bool {
		if res[i] == '<' {
			return true
		}
		if res[i] == '^' && res[j] != '<' {
			return true
		}
		if res[i] == 'v' && res[j] == '>' {
			return true
		}

		return false

	})
	ans := ""
	for _, st := range res {
		ans += string(st)
	}
	return ans
}

func sortkey(s string) string {

	res := []rune{}
	for _, el := range s {
		res = append(res, el)
	}
	// occ := occurences(res)
	sort.Slice(res, func(i, j int) bool {
		if res[i] == '^' && res[j] == '<' && s == "^<" {
			return true
		}
		if res[i] == '>' && res[j] == 'v' {
			return true
		}
		if res[i] == '>' && res[j] == '^' {
			return true
		}
		if res[i] == '<' && res[j] == 'v' {
			return true
		}
		return false

	})
	ans := ""
	for _, st := range res {
		ans += string(st)
	}
	return ans
}

func main() {
	start := time.Now()

	codes := strings.Split(inputDay, "\r\n")
	direc_grid := parseMatrix(direc)
	keypad_grid := parseMatrix(keypad)
	// fmt.Println(keypad)

	direc_a := find(direc_grid, 'A')
	keypad_a := find(keypad_grid, 'A')

	// robots := [4]Robot{direc_a, direc_a, keypad_a}

	fmt.Scanln()
	ans := 0
	// printMatrix(direc_grid)
	// printMatrix(keypad_grid)
	seen := map[[1]Robot]bool{}
	seen[[1]Robot{direc_a}] = true

	// a := solve2([1]Robot{direc_a}, 'v', direc_grid, 0, seen)
	// fmt.Println(a)
	// fmt.Println(try2(a, &[1]Robot{direc_a}, direc_grid))
	// fmt.Scanln()

	trans := map[State]string{}
	cells := "0123456789A"
	var pos2 = map[rune]Robot{}
	for _, cell := range cells {
		pos2[cell] = find(keypad_grid, cell)
	}
	a := ""
	for _, code := range codes {
		fmt.Println(code)
		robot := [1]Robot{keypad_a}
		res := ""
		for _, el := range code {
			// seen := map[[1]Robot]bool{}
			// seen[robot] = true
			a = ""
			verint, horint := pos2[el].x-robot[0].x, pos2[el].y-robot[0].y
			hor, ver := "", ""
			if verint > 0 {
				for range verint {
					ver += "v"
				}
			} else {
				for range -verint {
					ver += "^"
				}
			}
			if horint > 0 {
				for range horint {
					hor += ">"
				}
			} else {
				for range -horint {
					hor += "<"
				}
			}
			if robot[0].x == 3 && pos2[el].y == 0 {
				a += ver + hor
			} else if robot[0].y == 0 && pos2[el].x == 3 {
				a += hor + ver
			} else if verint < 0 {
				a += hor + ver
			} else {
				a += ver + hor
			}
			a += "A"
			trans[State{robot, el}] = a

			// fmt.Println(el, a)

			robot = [1]Robot{pos2[el]}
			res += a
			// fmt.Println(a, robots)

			// fmt.Println(a, robots)
		}

		// fmt.Println("first step found", code, res)

		cells := "^<>vA"
		pos := map[rune]Robot{}
		for _, cell := range cells {
			pos[cell] = find(direc_grid, cell)
		}

		a := ""

		for range 25 {
			res2 := ""
			robot := [1]Robot{direc_a}
			for _, el := range res {
				if trans[State{robot, el}] != "" {
					a = trans[State{robot, el}]
				} else {
					seen := map[[1]Robot]bool{}
					seen[robot] = true

					a = ""
					verint, horint := pos[el].x-robot[0].x, pos[el].y-robot[0].y
					hor, ver := "", ""
					if verint > 0 {
						for range verint {
							ver += "v"
						}
					} else {
						for range -verint {
							ver += "^"
						}
					}
					if horint > 0 {
						for range horint {
							hor += ">"
						}
					} else {
						for range -horint {
							hor += "<"
						}
					}
					if robot[0].y == 0 && pos[el].x == 0 {
						a += hor + ver
					} else if robot[0].x == 0 && pos[el].y == 0 {
						a += ver + hor
					} else if horint < 0 {
						a += hor + ver
					} else {
						a += ver + hor
					}
					a += "A"
					trans[State{robot, el}] = a
				}
				// fmt.Println(el, a)
				res2 += a
				robot = [1]Robot{pos[el]}

			}
			// fmt.Println(res)
			// fmt.Println(res2)
			res = res2
			try2(res2, &[1]Robot{direc_a}, direc_grid)
		}
		ans += numeric(code) * len(res)

	}
	fmt.Println(ans)
	fmt.Printf("Took %s", time.Since(start))
}
