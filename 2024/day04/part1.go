package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func cat(a, b, c, d byte) string {
	return string(a) + string(b) + string(c) + string(d)
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n")
	res := 0
	for i, line := range lines {
		for j, _ := range line {
			// horizontal check
			if j < len(line)-3 && line[j:j+4] == "XMAS" {
				res++
			}
			if j >= 3 && line[j-3:j+1] == "SAMX" {
				res++
			}

			// vertical check
			if i < len(lines)-3 && cat(lines[i][j], lines[i+1][j], lines[i+2][j], lines[i+3][j]) == "XMAS" {
				res++
			}
			if i >= 3 && cat(lines[i][j], lines[i-1][j], lines[i-2][j], lines[i-3][j]) == "XMAS" {
				res++
			}

			// diagonal check
			if (i >= 3 && j >= 3) && cat(lines[i][j], lines[i-1][j-1], lines[i-2][j-2], lines[i-3][j-3]) == "XMAS" {
				res++
			}
			if (i < len(lines)-3 && j < len(line)-3) && cat(lines[i][j], lines[i+1][j+1], lines[i+2][j+2], lines[i+3][j+3]) == "XMAS" {
				res++
			}
			if (i < len(lines)-3 && j >= 3) && cat(lines[i][j], lines[i+1][j-1], lines[i+2][j-2], lines[i+3][j-3]) == "XMAS" {
				res++
			}
			if (i >= 3 && j < len(line)-3) && cat(lines[i][j], lines[i-1][j+1], lines[i-2][j+2], lines[i-3][j+3]) == "XMAS" {
				res++
			}
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
