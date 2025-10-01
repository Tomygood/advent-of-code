package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed test.txt
var inputDay string

func nr(start int, rv int, ranges [][3]int) []int {

	for _, r := range ranges {
		if r[1] <= start && start < (r[1]+r[2]) {
			if r[1] <= start+rv && start+rv < (r[1]+r[2]) {
				return []int{r[0] + start - r[1], rv}
			} else {
				return append([]int{r[0] + start - r[1], r[2]}, nr(r[1]+r[2], start+rv-(r[1]+r[2]), ranges)...)
			}

		}
		// TODO
		return []int{}
	}

	// 15 11
	// 50 12 6

	return res
}

func next(cur []int, rans [](string)) []int {
	res := []int{}
	ranges := [][3]int{}
	for _, s := range rans {
		dest, source, am := 0, 0, 0
		fmt.Sscanf(s, "%d %d %d", &dest, &source, &am)
		ranges = append(ranges, [3]int{dest, source, am})
	}
	fmt.Println(ranges)
	for i := 0; i+1 < len(cur); i += 2 {
		starting_value := cur[i]
		range_value := cur[i+1]

		res = append(res, to_app)
	}
	return res
}

func main() {
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
	vals := strings.Split(strings.Split(steps[0], ": ")[1], " ")
	// fmt.Println(vals)
	values := []int{}
	// for i := 0; i < len(vals); i += 2 {
	// 	nv, r := 0, 0
	// 	fmt.Sscanf(vals[i], "%d", &nv)
	// 	fmt.Sscanf(vals[i+1], "%d", &r)
	// 	// fmt.Println("r", nv, "r")
	// 	for j := range r {
	// 		values = append(values, nv+j)
	// 	}
	// }
	for _, i := range vals {
		nv := 0
		fmt.Sscanf(i, "%d", &nv)
		// fmt.Println("r", nv, "r")

		values = append(values, nv)
	}
	fmt.Println(values)

	for _, step := range steps[1:] {
		ranges := strings.Split(step, "\n")[1:]
		values = next(values, ranges)
		fmt.Println(values)
	}
	// fmt.Println(values)
	min_val := values[0]
	for _, value := range values {
		if value < min_val {
			min_val = value
		}
	}
	fmt.Println(min_val)
}
