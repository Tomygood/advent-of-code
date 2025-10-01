// Shoutouts to Bar0lg
package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func concat(t []string) string {
	res := ""
	for _, a := range t {
		res += string(a)
	}
	return res
}

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func recalc(m map[int]int) int {
	res := 0
	for k, v := range m {
		res += power(8, k) * v
	}
	return res
}

func main() {
	start := time.Now()

	file := strings.Split(inputDay, "\r\n\r\n")
	prog := strings.Split(file[1], ": ")[1]

	commands := strings.Split(prog, ",")
	to_get := concat(commands)

	regs := map[string]int{}

	bit := len(to_get) - 1
	n := power(8, bit)
	atts := map[int]int{}

	for true {
		regs["A"] = n
		regs["B"], regs["C"] = 0, 0

		output := ""

		i := 0
		v := commands[0]
		for i+1 < len(commands) {
			v = commands[i]

			operand, _ := strconv.Atoi(commands[i+1])
			loperand := operand
			if operand == 4 {
				operand = regs["A"]
			} else if operand == 5 {
				operand = regs["B"]
			} else if operand == 6 {
				operand = regs["C"]
			}

			if v == "0" {
				regs["A"] = regs["A"] / power(2, operand)
			}
			if v == "1" {
				regs["B"] = regs["B"] ^ loperand
			}
			if v == "2" {
				regs["B"] = operand % 8
			}
			if v == "3" {
				if regs["A"] != 0 {
					i = loperand
					continue
				}
			}
			if v == "4" {
				regs["B"] = regs["B"] ^ regs["C"]
			}
			if v == "5" {
				output += strconv.Itoa(operand % 8)
			}
			if v == "6" {
				regs["B"] = regs["A"] / power(2, operand)
			}
			if v == "7" {
				regs["C"] = regs["A"] / power(2, operand)
			}
			i += 2
		}

		if len(output) >= 2 && output == to_get {
			break
		}
		if output[bit:] == to_get[bit:] {
			bit -= 1
		} else {
			n += power(8, bit)
			if atts[bit] == 7 {
				next := bit + 1
				for atts[next] == 7 {
					atts[next] = 0
					n = recalc(atts)
					next++
				}
				atts[bit] = 0
				atts[next] += 1
				n = recalc(atts)
				bit += 1
			} else {
				atts[bit]++
			}
		}
	}

	fmt.Println(n)
	fmt.Printf("Took %s", time.Since(start))
}
