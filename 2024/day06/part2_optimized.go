// solution using multithreading and various other optimizations (runs in just over 100 ms)

package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"
	"time"
)

//go:embed input.txt
var inputDay string

type Guard struct {
	x, y, dx, dy int
}

type Point struct {
	x, y int
}

func makeMatrix(n int) [][]byte {
	matrix := make([][]byte, n)
	for i := range matrix {
		matrix[i] = make([]byte, n)
	}
	return matrix
}

func parseMatrix(s string) [][]byte {
	var grid = strings.Split(s, "\r\n")
	mat := makeMatrix(len(grid))
	for i, line := range grid {
		for j := range line {
			mat[i][j] = line[j]
		}
	}
	return mat
}

func isLoop(obst Point, guard Guard, grid [][]byte) bool {
	seen := map[Guard]struct{}{}
	_, err := seen[guard]
	for err == false && guard.x >= 1 && guard.x < len(grid)-1 && guard.y >= 1 && guard.y < len(grid[0])-1 {
		nx, ny := guard.x+guard.dx, guard.y+guard.dy
		if grid[ny][nx] == '#' || (nx == obst.x && ny == obst.y) {
			if guard.dy == 0 {
				seen[guard] = struct{}{}
				guard.dx, guard.dy = 0, guard.dx
			} else {
				guard.dx, guard.dy = -guard.dy, 0
			}
		} else {
			guard.x, guard.y = nx, ny
		}
		_, err = seen[guard]
	}
	return err
}

func main() {
	start := time.Now()

	grid := parseMatrix(inputDay)
	guard := Guard{0, 0, 0, -1}
	for i, line := range grid {
		for j, c := range line {
			if c == '^' {
				guard.y = i
				guard.x = j
				grid[j][i] = '.'
			}
		}
	}

	guard2 := guard

	var wg sync.WaitGroup
	results := make(chan int, len(grid)*len(grid))

	num := map[[2]int]struct{}{}
	for guard.x >= 1 && guard.x < len(grid)-1 && guard.y >= 1 && guard.y < len(grid[0])-1 {
		nx, ny := guard.x+guard.dx, guard.y+guard.dy
		if grid[ny][nx] == '#' {
			if guard.dy == 0 {
				guard.dx, guard.dy = 0, guard.dx
			} else {
				guard.dx, guard.dy = -guard.dy, 0
			}
		} else {
			guard.x, guard.y = nx, ny
			_, err := num[[2]int{guard.x, guard.y}]
			if err == false {
				num[[2]int{guard.x, guard.y}] = struct{}{}
				wg.Add(1)
				go func(obst Point) {
					defer wg.Done()
					if isLoop(obst, guard2, grid) {
						results <- 1
					} else {
						results <- 0
					}
				}(Point{guard.x, guard.y})
			}
		}
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	res := 0
	for r := range results {
		res += r
	}
	fmt.Println(res)
	fmt.Printf("Took %s", time.Since(start))
}
