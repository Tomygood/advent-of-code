package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed input.txt
var inputDay string

func first_p(s []string, start int) (int, int) {
	for i := start; i < len(s); i++ {
		if s[i] == "." {
			res := 0
			j := i
			for j < len(s) && s[j] == "." {
				res++
				j++
			}
			return i, res
		}
	}
	return -1, -1
}

func last_non_p(s []string, end int) (int, int) {
	for i := end; i >= 0; i-- {
		if s[i] != "." {
			res := 0
			j := i
			for j < len(s) && s[j] == s[i] {
				if j == 0 {
					return -1, -1
				}
				res++
				j--
			}
			return j + 1, res
		}
	}
	return -1, -1
}

func main() {
	start := time.Now()

	full := []string{}
	for i := 0; i < len(inputDay); i += 2 {
		val, _ := strconv.Atoi(string(inputDay[i]))
		for k := range val {
			full, _ = append(full, strconv.Itoa(i/2)), k
		}

		if i+1 < len(inputDay) {
			val, _ = strconv.Atoi(string(inputDay[i+1]))
			for n := range val {
				full, _ = append(full, "."), n
			}
		}
	}

	in := make([]string, len(full))
	copy(in, full)

	cp, ps := first_p(full, 0)
	cn, ns := last_non_p(in, len(full)-1)

	for true {
		if cn <= 0 || cp == -1 {
			break
		}
		if ps >= ns && cp <= cn {
			for i := cp; i < cp+ns; i++ {
				full[i] = full[cn]
			}
			for j := cn; j < cn+ns; j++ {
				full[j] = "."
			}
			cn, ns = last_non_p(in, cn-1)
			cp, ps = first_p(full, 0)
		} else {
			if cp+ps < cn {
				cp, ps = first_p(full, cp+ps)
			} else {
				cn, ns = last_non_p(in, cn-1)
				cp, ps = first_p(full, 0)
			}
		}
	}

	res := 0
	for i, c := range full {
		if c != "." {
			val, _ := strconv.Atoi(c)
			res += val * i
		}
	}

	fmt.Println(res)
	fmt.Printf("Took %s", time.Since(start))
}
