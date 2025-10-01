// would take like 2 days to run but should work in theory (not cleaned)

package main

import (
	_ "embed"
	"fmt"
	"slices"
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

func solve(todo []string, swapsLeft int, swapped map[string]bool, swaps [4]string, seen map[[4]string]bool, res map[string]int, to_find map[string]int) string {
	if len(seen)%100_000 == 0 {
		fmt.Println(len(seen))
	}
	if swapsLeft == 0 {
		if calculate(todo, res, to_find) {
			return ""
		}
		return "XX"
	}
	for i, step := range todo {
		s := strings.Split(step, " ")
		gate1 := s[4]
		if !swapped[gate1] {
			swapped[gate1] = true
			for j, step2 := range todo {
				s2 := strings.Split(step2, " ")
				gate2 := s2[4]
				if gate2 != gate1 && !swapped[gate2] {
					curswap := ""
					if gate1 < gate2 {
						curswap = gate1 + gate2
					} else {
						curswap = gate2 + gate1
					}
					swaps[0] = curswap
					slices.Sort(swaps[:])
					if !seen[swaps] {
						seen[swaps] = true
						swapped[gate2] = true
						todo[i], todo[j] = todo[i][:len(todo[i])-3]+todo[j][len(todo[j])-3:], todo[j][:len(todo[j])-3]+todo[i][len(todo[i])-3:]
						n := solve(todo, swapsLeft-1, swapped, swaps, seen, res, to_find)
						if len(n) == 0 || n[0] != 'X' {
							fmt.Println(todo)

							return n + gate1 + "," + gate2 + ","
						}
						todo[i], todo[j] = todo[i][:len(todo[i])-3]+todo[j][len(todo[j])-3:], todo[j][:len(todo[j])-3]+todo[i][len(todo[i])-3:]
						swapped[gate2] = false
					}
					for k, swap := range swaps {
						if len(swap) > 0 && (swap == curswap) {
							swaps[k] = ""
							break
						}
					}

				}
			}
			swapped[gate1] = false
		}
	}
	return "XX"
}

func calculate(todo []string, res map[string]int, to_find map[string]int) bool {

	nres := make(map[string]int)
	for k, v := range res {
		nres[k] = v
	}
	step := todo[0]
	depth := 0
	for len(todo) > 0 {
		if depth == 4000 {
			break
		}
		step = todo[0]
		s := strings.Split(step, " ")
		gate1, op, gate2, ngate := s[0], s[1], s[2], s[4]
		if is_in(gate1, nres) && is_in(gate2, nres) {
			val1, val2 := nres[gate1], nres[gate2]
			nres[ngate] = result(op, val1, val2)
			if ngate[0] == 'z' {
				if nres[ngate] != to_find[ngate] {
					return false
				}
			}
			todo = todo[1:]
		} else {
			todo = append(todo[1:], todo[0])
		}
		depth++
	}
	if depth == 4000 {
		return false
	}
	return true
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
	fmt.Println(xb)
	zToGetI := strconv.FormatInt(xb+yb, 2)

	zToGetM := map[string]int{}
	for i := range len(zToGetI) {
		if len(zToGetI)-i-1 >= 10 {
			zToGetM["z"+strconv.Itoa(len(zToGetI)-i-1)] = atoi(string(zToGetI[i]))
		} else {
			zToGetM["z0"+strconv.Itoa(len(zToGetI)-i-1)] = atoi(string(zToGetI[i]))

		}
	}
	fmt.Println(zToGetI)
	fmt.Println(zToGetM)

	ans := solve(todo, 4, map[string]bool{}, [4]string{}, map[[4]string]bool{}, res, zToGetM)
	nm := strings.Split(ans, ",")
	fmt.Println(nm)
	slices.Sort(nm)

	for i, v := range nm[1:] {
		if i == 0 {
			print(v)
		} else {
			print("," + v)
		}
	}
	fmt.Println("")
	fmt.Printf("Took %s", time.Since(start))
}
