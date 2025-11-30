package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
)

// 2D deltas
var deltas = [4]Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// 3D deltas
var deltas3D = [6]Point3D{{0, 0, 1}, {0, 0, -1}, {0, 1, 0}, {0, -1, 0}, {1, 0, 0}, {-1, 0, 0}}

// Abs(x) returns the absolute value of x as an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Atoi performs strconv.Atoi and ignores the error return argument
func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// Power(a, b) returns a^b as an integer
func Power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// Mod performs the modulo operation and always returns positive values
func Mod(a, b int) int {
	return (a%b + b) % b
}

// Concat turns a slice of runes into a string
func Concat(t []rune) string {
	var res string
	for _, a := range t {
		res += string(a)
	}
	return res
}

// ConcatInts concatenates two integers
func ConcatInts(a, b int) int {
	return Atoi(strconv.Itoa(a) + strconv.Itoa(b))
}

// ConcatIntSlice concatenates a slice of integers
func ConcatIntSlice(l []int) int {
	var res int
	for _, a := range l {
		res = ConcatInts(res, a)
	}
	return res
}

// Lines returns a slices containing each line of the input file
func Lines(s string) []string {
	return strings.Split(s, "\r\n")
}

// Generic 2D point structure
type Point struct {
	x, y int
}

// Generic 2D point structure with directions dx and dy
type PointD struct {
	x, y, dx, dy int
}

// Generic 3D point structure
type Point3D struct {
	x, y, z int
}

// Generic 3D point structure with directions dx, dy and dz
type Point3DD struct {
	x, y, z, dx, dy, dz int
}

// ToClipboard writes argument into user’s clipboard
func ToClipboard[E any](res E) {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	s := fmt.Sprint(res)

	clipboard.Write(clipboard.FmtText, []byte(s))
}
