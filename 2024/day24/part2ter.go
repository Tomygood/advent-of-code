// used only to display info, doesn’t solve the problem on its own (not cleaned either)

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

func reverse(l []string) []string {
	res := []string{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

func is_in(val string, m map[string]int) bool {
	_, ok := m[val]
	return ok
}

func mtb(start string, m map[string]int) string {
	z, ans := "", ""
	for i := 99; i >= 0; i-- {
		if i < 10 {
			z = start + "0" + strconv.Itoa(i)
		} else {
			z = start + strconv.Itoa(i)
		}
		if _, ok := m[z]; ok {
			ans += strconv.Itoa(m[z])
		}
	}
	return ans
}

func mtbmap(start string, m map[string]int) map[string]int {
	z, ans := "", map[string]int{}
	for i := 99; i >= 0; i-- {
		if i < 10 {
			z = start + "0" + strconv.Itoa(i)
		} else {
			z = start + strconv.Itoa(i)
		}
		if _, ok := m[z]; ok {
			ans[z] = m[z]
		}
	}
	return ans
}

func noDupes[L ~[]E, E comparable](l L) []E {
	seen := map[E]bool{}
	res := []E{}
	for _, el := range l {
		if !seen[el] {
			res = append(res, el)
			seen[el] = true
		}
	}

	return res
}

func calculate(todo []string, res map[string]int) string {
	nres := make(map[string]int)
	for k, v := range res {
		nres[k] = v
	}
	step := todo[0]

	for len(todo) > 0 {
		fmt.Println(todo)
		step = todo[0]
		s := strings.Split(step, " ")
		gate1, op, gate2, ngate := s[0], s[1], s[2], s[4]
		if is_in(gate1, nres) && is_in(gate2, nres) {
			val1, val2 := nres[gate1], nres[gate2]
			nres[ngate] = result(op, val1, val2)
			todo = todo[1:]
		} else {
			todo = append(todo[1:], todo[0])
		}
	}
	return mtb("z", nres)
}

func endof(todo []string, tofind string) int {
	for i, step := range todo {
		if step[len(step)-3:] == tofind {
			return i
		}
	}
	return 0
}

func beginningof(todo []string, tofind string) int {
	for _, step := range todo {
		if step[1:3] == tofind {
			if step[4] == 'A' {
				fmt.Println("Retenue : ", step[len(step)-3:])
			}
			if step[4] == 'X' {
				fmt.Println("Rés. : ", step[len(step)-3:])
			}

		}
	}
	return 0
}

func main() {
	start := time.Now()

	lines := strings.Split(inputDay, "\r\n\r\n")

	res := map[string]int{}

	for _, starting := range strings.Split(lines[0], "\r\n") {
		s := strings.Split(starting, ": ")
		res[s[0]] = atoi(s[1])
	}
	zs := map[string]int{}
	todo := strings.Split(lines[1], "\r\n")

	for _, step := range todo {
		s := strings.Split(step, " ")
		gate1, op, gate2 := s[0], s[1], s[2]
		if gate1[0] == 'x' && gate2[0] == 'y' && gate1[1:] == gate2[1:] {
			zs["z"+gate1[1:]] = result(op, res[gate1], res[gate2])
		}
	}

	// zToGet := mtbmap("z", zs)
	xb, _ := strconv.ParseInt(mtb("x", res), 2, 64)
	yb, _ := strconv.ParseInt(mtb("y", res), 2, 64)
	zToGetI := strconv.FormatInt(xb+yb, 2)

	zToGetM := map[string]int{}
	for i := range len(zToGetI) {
		if len(zToGetI)-i-1 >= 10 {
			zToGetM["z"+strconv.Itoa(len(zToGetI)-i-1)] = atoi(string(zToGetI[i]))
		} else {
			zToGetM["z0"+strconv.Itoa(len(zToGetI)-i-1)] = atoi(string(zToGetI[i]))

		}
	}

	zInit := calculate(todo, res)

	fmt.Println(zInit)
	fmt.Println(zToGetI)

	carries := map[string]string{}
	j := ""
	for i, _ := range zInit {
		if i < 10 {
			j = "0" + strconv.Itoa(i)
		} else {
			j = strconv.Itoa(i)
		}
		beginningof(todo, j)
		fmt.Println(todo[endof(todo, "z"+j)])
	}

	fmt.Println(carries)
	fmt.Printf("Took %s", time.Since(start))
}
