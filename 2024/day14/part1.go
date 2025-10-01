package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	start := time.Now()

	width := 101
	height := 103

	lines := strings.Split(inputDay, "\r\n")
	nw, ne, sw, se := 0, 0, 0, 0

	for _, robot := range lines {
		px, py, vx, vy := 0, 0, 0, 0
		fmt.Sscanf(robot, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		for range 100 {
			px = mod((px + vx), width)
			py = mod((py + vy), height)
		}

		if px < width/2 && py < height/2 {
			nw++
		} else if px < width/2 && py > height/2 {
			sw++
		} else if px > width/2 && py < height/2 {
			ne++
		} else if px > width/2 && py > height/2 {
			se++
		}
	}

	fmt.Println(nw * ne * sw * se)

	fmt.Printf("Took %s", time.Since(start))
}
