package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func reach(px, py, xa, ya, xb, yb, moves int) int {
	m := 0
	a, b := 0, 0

	for i := 0; xa*i <= px && ya*i <= py; i++ {
		for j := 0; xa*i+xb*j <= px && ya*i+yb*j <= py; j++ {
			if xa*i+xb*j == px && ya*i+yb*j == py {
				if -(i + j) < m {
					a, b = i, j
					m = -(i + j)
				}
			}
		}
	}
	return 3*a + b

}

func parse(s string) (int, int) {
	l := strings.Split(strings.Split(s, ": ")[1], ", ")
	X, Y := 0, 0
	fmt.Sscanf(l[0], "X+%d", &X)
	fmt.Sscanf(l[1], "Y+%d", &Y)
	return X, Y

}

func main() {
	start := time.Now()
	res := 0

	for _, game := range strings.Split(inputDay, "\r\n\r\n") {
		lines := strings.Split(game, "\r\n")

		Xa, Ya := parse(lines[0])
		Xb, Yb := parse(lines[1])

		l := strings.Split(strings.Split(lines[2], ": ")[1], ", ")
		pX, pY := 0, 0
		fmt.Sscanf(l[0], "X=%d", &pX)
		fmt.Sscanf(l[1], "Y=%d", &pY)
		res += reach(pX, pY, Xa, Ya, Xb, Yb, 0)
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))

}
