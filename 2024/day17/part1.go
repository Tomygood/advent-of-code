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

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	start := time.Now()

	file := strings.Split(inputDay, "\r\n\r\n")
	registers := strings.Split(file[0], "\r\n")
	prog := strings.Split(file[1], ": ")[1]

	output := ""
	regs := make(map[string]int)
	for _, reg := range registers {
		ind := ""
		fmt.Sscanf(reg, "Register %s", &ind)
		val, _ := strconv.Atoi(strings.Split(reg, ": ")[1])
		regs[string(ind[0])] = val
	}

	commands := strings.Split(prog, ",")
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
			regs["A"] = int(float64(regs["A"]) / float64(power(2, operand)))
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
			output += strconv.Itoa(operand%8) + ","
		}
		if v == "6" {
			regs["B"] = regs["A"] / power(2, operand)
		}
		if v == "7" {
			regs["C"] = regs["A"] / power(2, operand)
		}
		i += 2
	}

	fmt.Println(output[:len(output)-1])
	fmt.Printf("Took %s", time.Since(start))
}
