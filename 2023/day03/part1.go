package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputDay string

func main() {
	lines := strings.Split(inputDay, "\n")
	total := 0
	for k, line := range lines {
		start := -1
		for i, c := range line {
			if c-48 >= 0 && c-48 <= 9 && start == -1 {
				start = i
			}
			if (c-48 < 0 || c-48 > 9 || i == len(line)-1) && start != -1 {
				// check if number if valid
				valid := false
				if (start != 0 && line[start-1] != '.') || (i != len(line)-1 && c != '.') {
					valid = true
				}
				for j := start - 1; j <= i; j++ {
					if (k != 0 && (start != 0 || j != start-1) && (i != len(line)-1 || j != i) && lines[k-1][j] != '.') || (k != len(lines)-1 && (start != 0 || j != start-1) && (i != len(line)-1 || j != i) && lines[k+1][j] != '.') {
						valid = true
					}
				}
				if valid {
					num := 0
					fmt.Sscanf(line[start:i], "%d", &num)
					total += num
				}
				start = -1
			}

		}
	}
	fmt.Println(total)
}
