package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n")

	connections := map[string]map[string]bool{}

	for _, line := range lines {
		s := strings.Split(line, "-")
		a, b := s[0], s[1]
		if connections[a] == nil {
			connections[a] = map[string]bool{}
		}
		if connections[b] == nil {
			connections[b] = map[string]bool{}
		}
		connections[a][b] = true
		connections[b][a] = true
	}

	seen := map[[3]string]bool{}

	for a, am := range connections {
		for b := range am {
			for c := range connections[b] {
				if c != a && connections[a][c] {
					sl := []string{a, b, c}
					slices.Sort(sl)
					arr := [3]string{sl[0], sl[1], sl[2]}
					seen[arr] = true
				}
			}
		}
	}
	res := 0

	for k := range seen {
		for _, com := range k {
			if com[0] == 't' {
				res++
				break
			}
		}
	}

	fmt.Println(res)

	fmt.Printf("Took %s", time.Since(start))
}
