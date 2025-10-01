package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func reach(px, py, xa, ya, xb, yb float64) int {
	a := int(math.Round(px/xa - xb/xa*(py-(ya*(px/xa)))/(yb-(ya*(xb/xa)))))
	b := int(math.Round((py - (ya * (px / xa))) / (yb - (ya * (xb / xa)))))

	if int(xa)*a+int(xb)*b == int(px) && int(ya)*a+int(yb)*b == int(py) {
		return 3*a + b
	}
	return 0
}

func parse(s string) (float64, float64) {
	l := strings.Split(strings.Split(s, ": ")[1], ", ")
	X, Y := 0., 0.
	fmt.Sscanf(l[0], "X+%f", &X)
	fmt.Sscanf(l[1], "Y+%f", &Y)
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
		pX, pY := 0., 0.
		fmt.Sscanf(l[0], "X=%f", &pX)
		fmt.Sscanf(l[1], "Y=%f", &pY)

		res += reach(pX+10000000000000, pY+10000000000000, Xa, Ya, Xb, Yb)
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))

}
