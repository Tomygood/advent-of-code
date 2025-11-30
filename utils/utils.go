package utils

import (
	"math"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Concat(t []rune) string {
	res := ""
	for _, a := range t {
		res += string(a)
	}
	return res
}

func ListToString(s []rune) string {
	res := ""
	for _, c := range s {
		res += string(c)
	}
	return res
}

func ConcatInts(a, b int) int {
	res, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return res
}

type Point struct {
	x, y int
}

type PointD struct {
	x, y, dx, dy int
}

type Point3D struct {
	x, y, z int
}

type Point3DD struct {
	x, y, z, dx, dy, dz int
}

// deltas := [4]Point{Point{0, 1}, Point{0, -1}, Point{1, 0}, Point{-1, 0}}
