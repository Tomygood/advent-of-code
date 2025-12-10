package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Tomygood/advent-of-code/utils"
)

func createKey(s []bool) string { return fmt.Sprint(s) }

func process(buttons [][]int, goal []bool) int {
	init := []bool{}
	for range len(goal) {
		init = append(init, false)
	}
	seen := map[string]bool{createKey(init): true}

	return process_rec(init, buttons, goal, seen)
}

func process_rec(state []bool, buttons [][]int, goal []bool, seen map[string]bool) int {
	if slices.Equal(state, goal) {
		return 0
	}

	l := math.MaxInt - 1
	for _, button := range buttons {

		for _, toggle := range button {
			state[toggle] = !state[toggle]
		}

		if seen[createKey(state)] {
			for _, toggle := range button {
				state[toggle] = !state[toggle]
			}
			break
		}
		seen[createKey(state)] = true

		s := process_rec(state, buttons, goal, seen)
		if s < l {
			l = s
		}

		seen[createKey(state)] = false
		for _, toggle := range button {
			state[toggle] = !state[toggle]
		}
	}

	return l + 1
}

func part1() {

	machines := utils.Lines(inputDay)

	res := 0

	for _, machine := range machines {
		a := strings.Split(machine, " ")
		g := a[0]
		goal := []bool{}
		for i := 1; i < len(g)-1; i++ {
			goal = append(goal, g[i] == '#')
		}

		buttons := [][]int{}
		for _, button := range a[1 : len(a)-1] {
			toadd := []int{}
			for _, toggle := range strings.Split(button[1:len(button)-1], ",") {
				toadd = append(toadd, utils.Atoi(toggle))
			}
			buttons = append(buttons, toadd)
		}
		res += process(buttons, goal)
	}

	fmt.Println(res)
	utils.ToClipboard(res)
}
