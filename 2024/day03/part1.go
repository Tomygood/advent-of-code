package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var inputDay1 string

func main() {
	start := time.Now()

	res := 0
	for i := range len(inputDay1) - 7 {
		n1, n2 := 0, 0
		_, err := fmt.Sscanf(inputDay1[i:], "mul(%d,%d)", &n1, &n2)
		if err == nil {
			res += n1 * n2
			i += 7
		}

	}
	fmt.Println(res)

	fmt.Println(time.Since(start))
}
