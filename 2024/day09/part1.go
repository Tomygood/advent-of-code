package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed input.txt
var inputDay string

func first_p(s []string, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == "." {
			return i
		}
	}
	return -1
}

func last_non_p(s []string, end int) int {
	for i := end; i >= 0; i-- {
		if s[i] != "." {
			return i
		}
	}
	return -1
}

func main() {
	start := time.Now()

	full := []string{}
	for i := 0; i < len(inputDay); i += 2 {
		val, _ := strconv.Atoi(string(inputDay[i]))
		for k := range val {
			full, _ = append(full, strconv.Itoa(i/2)), k
		}

		if i+1 < len(inputDay) {
			val, _ = strconv.Atoi(string(inputDay[i+1]))
			for n := range val {
				full, _ = append(full, "."), n
			}
		}
	}

	current := first_p(full, 0)
	ts := last_non_p(full, len(full)-1)
	for current != ts+1 {
		full[ts], full[current] = ".", full[ts]
		current = first_p(full, current+1)
		ts = last_non_p(full, ts-1)
	}

	res := 0
	for i, c := range full {
		if c != "." {
			val, _ := strconv.Atoi(c)
			res += val * i
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
