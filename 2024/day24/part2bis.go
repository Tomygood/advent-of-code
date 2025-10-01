// would take like 18 hours to run but should work in theory (also not cleaned)

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

func calculate_keep(todo []string, res map[string]int) (string, []int) {
	nres := make(map[string]int)
	for k, v := range res {
		nres[k] = v
	}
	step := todo[0]
	prouiy := []int{}
	for i := range len(todo) {
		prouiy = append(prouiy, i)
	}
	order := []int{}

	for len(todo) > 0 {
		step = todo[0]
		s := strings.Split(step, " ")
		gate1, op, gate2, ngate := s[0], s[1], s[2], s[4]
		if is_in(gate1, nres) && is_in(gate2, nres) {
			val1, val2 := nres[gate1], nres[gate2]
			nres[ngate] = result(op, val1, val2)
			order = append(order, prouiy[0])
			todo = todo[1:]
			prouiy = prouiy[1:]
		} else {
			todo = append(todo[1:], todo[0])
			prouiy = append(prouiy[1:], prouiy[0])

		}
	}
	return mtb("z", nres), order
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

func solve(todo []string, suspects []int, swapsLeft int, swapped map[string]bool, swaps [4]string, seen map[[4]string]bool, res map[string]int, to_find map[string]int) string {
	if len(seen)%100_000 == 0 {
		fmt.Println(len(seen))
	}
	if swapsLeft == 2 {
		fmt.Println(swaps)
	}
	if swapsLeft == 0 {
		if calculate(todo, res, to_find) {
			return ""
		}
		return "XX"
	}
	for _, i := range suspects {
		s := strings.Split(todo[i], " ")
		gate1 := s[4]
		if !swapped[gate1] {
			swapped[gate1] = true
			for _, j := range suspects {
				s2 := strings.Split(todo[j], " ")
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
						n := solve(todo, suspects, swapsLeft-1, swapped, swaps, seen, res, to_find)
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

func noDupes(l []int) []int {
	seen := map[int]bool{}
	res := []int{}
	for _, el := range l {
		if !seen[el] {
			res = append(res, el)
			seen[el] = true
		}
	}
	return res
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

	zInit, order := calculate_keep(todo, res)

	ntodo := []string{}
	for i := range order {
		ntodo = append(ntodo, todo[order[i]])
	}

	suspects := []int{}
	next := []string{}
	for i, c := range zInit {
		if c != rune(zToGetI[i]) {
			next = append(next, "z"+strconv.Itoa(len(zInit)-i-1))
		}
	}

	fmt.Println(next)
	for len(next) > 0 {
		nnext := []string{}

		for _, child := range next {
			for i, step := range ntodo {
				s := strings.Split(step, " ")
				gate1, gate2, gate3 := s[0], s[2], s[4]
				if gate3 == child {
					suspects = append(suspects, i)
					if gate1[0] != 'x' && gate1[0] != 'y' {
						nnext = append(nnext, gate1)
					}
					if gate2[0] != 'x' && gate2[0] != 'y' {
						nnext = append(nnext, gate2)
					}
				}
			}
		}
		next = nnext
	}
	suspects = noDupes(suspects)

	fmt.Println(zInit)
	fmt.Println(zToGetI)

	fmt.Println(suspects)
	for _, sus := range suspects {
		fmt.Println(ntodo[sus])
	}
	// fmt.Println(ntodo)

	ans := solve(ntodo, suspects, 4, map[string]bool{}, [4]string{}, map[[4]string]bool{}, res, zToGetM)
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
