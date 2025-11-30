// replace any with desired type
package utils

import "slices"

// NoDupes returns a version of the slice without duplicate elements
func NoDupes[L ~[]E, E comparable](l L) []E {
	seen := map[E]bool{}
	res := []E{}
	for _, el := range l {
		if !seen[el] {
			res = append(res, el)
			seen[el] = true
		}
	}

	return res
}

// Flatten concatenates a slice of same-type slices
func Flatten[L ~[][]E, E any](lists L) []E {
	res := []E{}
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

// Union performs the union operator on two slices
func Union[L ~[]E, E comparable](l1, l2 L) []E {
	return NoDupes(Flatten([][]E{l1, l2}))
}

// Reverse returns a reversed version of the slice given
func Reverse[L ~[]E, E any](l L) []E {
	res := []E{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

// Occurences returns a map counting the amount of each element in the slice
func Occurences[L ~[]E, E comparable](l L) map[E]int {
	res := map[E]int{}
	for _, el := range l {
		res[el]++
	}
	return res
}

// Intersection performs the intersection operator on two slices
func Intersection[L ~[]E, E comparable](l1, l2 L) []E {
	res := []E{}
	for _, el := range l1 {
		if slices.Contains(l2, el) {
			res = append(res, el)
		}
	}
	return res
}

// Permutations returns a slice containing all possible permutations of the slice given
func Permutations[L ~[]E, E any](l L) [][]E {
	if len(l) <= 1 {
		return [][]E{l}
	}
	res := [][]E{}
	perms := Permutations(l[1:])
	for _, perm := range perms {
		for i := 0; i < len(perm)+1; i++ {
			newPerm := append(append([]E{}, perm[:i]...), l[0])
			newPerm = append(newPerm, perm[i:]...)
			res = append(res, newPerm)
		}
	}

	return res
}
