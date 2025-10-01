package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func cat(a, b byte) string {
	return string(a) + string(b)
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n")
	res := 0
	for i, line := range lines {
		for j, c := range line {
			if i >= 1 && i < len(lines)-1 && j >= 1 && j < len(line)-1 && string(c) == "A" {
				diag1 := cat(lines[i-1][j-1], lines[i+1][j+1])
				diag2 := cat(lines[i-1][j+1], lines[i+1][j-1])
				if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
					res++
				}
			}
		}
	}
	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
