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
	seen := map[[20]string]bool{}

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
		seen[[20]string{a}] = true
	}

	for true {
		nseen := make(map[[20]string]bool)
		for party := range seen {
			nparty := []string{}
			for _, participant := range party {
				if len(participant) > 1 {
					nparty = append(nparty, participant)
				}
			}

			for _, participant := range party {
				if len(participant) > 1 {
					for partconn := range connections[participant] {
						canEnter := true
						for _, neigh := range party {
							if len(neigh) > 1 && !connections[neigh][partconn] {
								canEnter = false
							}
						}
						if canEnter {
							newparty := make([]string, len(nparty))
							copy(newparty, nparty)
							newparty = append(newparty, partconn)
							slices.Sort(newparty)
							finalparty := [20]string{}
							for i, part := range newparty {
								finalparty[i] = part
							}
							nseen[finalparty] = true
						}
					}
				}
			}
		}
		if len(nseen) == 0 {
			break
		}
		seen = nseen
	}

	for k := range seen {
		for i, part := range k {
			if len(part) > 1 {
				if i != 0 {
					print(",", part)
				} else {
					print(part)
				}
			}
		}
	}
	fmt.Println("")

	fmt.Printf("Took %s", time.Since(start))
}
