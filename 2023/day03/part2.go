package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

type Gear struct {
	x, y  int
	part1 int
	part2 int
}

func find_gear(gears []Gear, x int, y int) int {
	for i, gear := range gears {
		if gear.x == x && gear.y == y {
			return i
		}
	}
	return -1
}

func add_part(g *Gear, n int) {
	if g.part1 == 0 {
		g.part1 = n
	} else {
		g.part2 = n
	}
}

func main() {
	lines := strings.Split(inputDay, "\n")
	total := 0
	gears := []Gear{}
	for k, line := range lines {
		for i, c := range line {
			if c == '*' {
				gears = append(gears, Gear{k, i, 0, 0})
			}
		}
	}
	for k, line := range lines {
		start := -1
		for i, c := range line {
			if c-48 >= 0 && c-48 <= 9 && start == -1 {
				start = i
			}
			if (c-48 < 0 || c-48 > 9 || i == len(line)-1) && start != -1 {

				num := 0
				g := -1
				fmt.Sscanf(line[start:i], "%d", &num)

				if start != 0 && line[start-1] == '*' {
					g = find_gear(gears, k, start-1)
				}
				if i != len(line)-1 && c == '*' {
					g = find_gear(gears, k, i)
				}
				for j := start - 1; j <= i; j++ {
					if k != 0 && (start != 0 || j != start-1) && (i != len(line)-1 || j != i) && lines[k-1][j] == '*' {
						g = find_gear(gears, k-1, j)
					}
					if k != len(lines)-1 && (start != 0 || j != start-1) && (i != len(line)-1 || j != i) && lines[k+1][j] == '*' {
						g = find_gear(gears, k+1, j)
					}
				}
				if g != -1 {
					add_part(&gears[g], num)
				}

				start = -1
			}

		}
	}
	for _, gear := range gears {
		total += gear.part1 * gear.part2
	}
	fmt.Println(total)
}
