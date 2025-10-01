package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func result(op string, gate1, gate2 int) int {
	if op == "AND" {
		return Btoi(gate1 == 1 && gate2 == 1)
	}
	if op == "OR" {
		return Btoi(gate1 == 1 || gate2 == 1)
	} else {
		return Btoi(gate1 != gate2)
	}
}

func is_in(val string, m map[string]int) bool {
	_, ok := m[val]
	return ok
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")

	res := map[string]int{}

	for _, starting := range strings.Split(lines[0], "\r\n") {
		s := strings.Split(starting, ": ")
		res[s[0]] = atoi(s[1])
	}

	zs := []string{}
	todo := strings.Split(lines[1], "\r\n")
	step := todo[0]

	for len(todo) > 0 {
		step = todo[0]
		s := strings.Split(step, " ")
		gate1, op, gate2, ngate := s[0], s[1], s[2], s[4]
		if is_in(gate1, res) && is_in(gate2, res) {
			val1, val2 := res[gate1], res[gate2]
			res[ngate] = result(op, val1, val2)
			if ngate[0] == 'z' {
				zs = append(zs, ngate)
			}
			todo = todo[1:]
		} else {
			todo = append(todo[1:], todo[0])
		}
	}

	ans := ""
	z := ""
	for i := 99; i >= 0; i-- {
		if i < 10 {
			z = "z0" + strconv.Itoa(i)
		} else {
			z = "z" + strconv.Itoa(i)
		}
		if _, ok := res[z]; ok {
			ans += strconv.Itoa(res[z])
		}
	}

	v, _ := strconv.ParseInt(ans, 2, 64)
	fmt.Println(v)

	fmt.Printf("Took %s", time.Since(start))
}
