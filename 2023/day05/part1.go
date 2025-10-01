package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed test.txt
var inputDay string

func next(cur []int, rans [](string)) []int {
	res := []int{}
	for _, val := range cur {
		to_app := val
		for _, s := range rans {
			dest, source, am := 0, 0, 0
			fmt.Sscanf(s, "%d %d %d", &dest, &source, &am)
			if source <= val && val < (source+am) {
				to_app = dest + val - source
			}
		}
		res = append(res, to_app)
	}
	return res
}

func main() {
	// hum… don’t bother understanding this I learnt too late that Windows is stupid
	re := regexp.MustCompile(`(?m)^[a-zA-Z].*`)
	headers := re.FindAllStringIndex(inputDay, -1)
	var steps []string
	for i, header := range headers {
		start := header[0]
		var end int
		if i < len(headers)-1 {
			end = headers[i+1][0]
		} else {
			end = len(inputDay)
		}
		steps = append(steps, strings.TrimSpace(inputDay[start:end]))
	}
	// yeah…
	vals := strings.Split(strings.Split(steps[0], ": ")[1], " ")
	fmt.Println(vals)
	values := []int{}
	for _, i := range vals {
		nv := 0
		fmt.Sscanf(i, "%d", &nv)
		values = append(values, nv)
	}

	for _, step := range steps[1:] {
		ranges := strings.Split(step, "\n")[1:]
		values = next(values, ranges)
		fmt.Println(values)

	}
	fmt.Println(values)
	min_val := values[0]
	for _, value := range values {
		if value < min_val {
			min_val = value
		}
	}
	fmt.Println(min_val)
}
