package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
)

// 2D deltas
var Deltas = [4]Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// 2D 8-way deltas
var DeltasOrtho = [8]Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

// 3D deltas
var Deltas3D = [6]Point3D{{0, 0, 1}, {0, 0, -1}, {0, 1, 0}, {0, -1, 0}, {1, 0, 0}, {-1, 0, 0}}

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

// ConcatInts concatenates multiple integers
func ConcatInts(a ...int) int {
	var res string
	for _, e := range a {
		res += strconv.Itoa(e)
	}
	return Atoi(res)
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
	lines := strings.Split(s, "\r\n")
	var res []string
	for _, line := range lines {
		if line != "" {
			res = append(res, line)
		}
	}

	return res
}

// Generic 2D point structure
type Point struct {
	X, Y int
}

// Generic 2D point structure with directions dx and dy
type PointD struct {
	X, Y, Dx, Dy int
}

// Generic 3D point structure
type Point3D struct {
	X, Y, Z int
}

// Generic 3D point structure with directions dx, dy and dz
type Point3DD struct {
	X, Y, Z, Dx, Dy, Dz int
}

// Returns the sum of two points
func AddPoints(p1, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

// MovePoint moves a point to its next value using its current direction
func MovePoint(p PointD) PointD {
	return PointD{p.X + p.Dx, p.Y + p.Dy, p.Dx, p.Dy}
}

// ToClipboard writes argument into system clipboard
func ToClipboard[E any](res E) {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	s := fmt.Sprint(res)

	clipboard.Write(clipboard.FmtText, []byte(s))
}
