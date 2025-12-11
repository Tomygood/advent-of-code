// SEE part2.py

package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func createKey2(s []int) string { return fmt.Sprint(s) }

func process2(buttons [][]int, goal []int) int {
	init := []int{}
	for range len(goal) {
		init = append(init, 0)
	}
	seen := map[string]int{}

	return process_rec2(init, buttons, goal, seen, 0)
}

func process_rec2(state []int, buttons [][]int, goal []int, seen map[string]int, depth int) int {
	//fmt.Println(state, seen, depth)
	//fmt.Scanln()
	if slices.Equal(state, goal) {
		fmt.Println("end")
		seen[createKey2(state)] = 0
		return 0
	}

	l := math.MaxInt
	for _, button := range buttons {

		new_state := slices.Clone(state)

		for _, toggle := range button {
			new_state[toggle]++
		}

		var s int
		if seen[createKey2(new_state)] > 0 {
			s = seen[createKey2(new_state)]
		} else {
			flag := false
			for i, c := range goal {
				if new_state[i] > c {
					flag = true
					break
				}
			}
			if flag {
				break
			}

			s = process_rec2(new_state, buttons, goal, seen, depth+1)
		}
		if s < l {
			l = s
		}
	}
	if l == math.MaxInt {
		return l
	}
	seen[createKey2(state)] = l + 1
	return l + 1
}

func part2() {

	machines := utils.Lines(inputDay)

	res := 0

	for i, machine := range machines {
		fmt.Println(i)
		a := strings.Split(machine, " ")
		g := a[len(a)-1]
		goal := []int{}
		for _, counter := range strings.Split(g[1:len(g)-1], ",") {
			goal = append(goal, utils.Atoi(counter))
		}

		//fmt.Println(goal)

		buttons := [][]int{}
		for _, button := range a[1 : len(a)-1] {
			toadd := []int{}
			for _, toggle := range strings.Split(button[1:len(button)-1], ",") {
				toadd = append(toadd, utils.Atoi(toggle))
			}
			buttons = append(buttons, toadd)
		}

		res += process2(buttons, goal)

	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
