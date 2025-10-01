package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var inputDay2 string

func main() {
	start := time.Now()

	res := 0
	do := true
	for i := range len(inputDay2) - 7 {
		cur := inputDay2[i:]
		if cur[:4] == "do()" {
			do = true
			i += 3
		} else {
			if cur[:7] == "don't()" {
				do = false
				i += 6
			} else {
				if do {
					n1, n2 := 0, 0
					_, err := fmt.Sscanf(cur, "mul(%d,%d)", &n1, &n2)
					if err == nil {
						res += n1 * n2
						i += 7
					}
				}
			}
		}
	}
	fmt.Println(res)

	fmt.Println(time.Since(start))
}
