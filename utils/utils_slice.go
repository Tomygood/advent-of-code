// replace any with desired type
package utils

import "slices"

// noDupes returns a version of the slice without duplicate elements
func noDupes[L ~[]E, E comparable](l L) []E {
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

// flatten concatenates a slice of same-type slices
func flatten[L ~[][]E, E any](lists L) []E {
	res := []E{}
	for _, list := range lists {
		for _, el := range list {
			res = append(res, el)
		}
	}
	return res
}

// union performs the union operator on two slices
func union[L ~[]E, E comparable](l1, l2 L) []E {
	return noDupes(flatten([][]E{l1, l2}))
}

// reverse returns a reversed version of the slice given
func reverse[L ~[]E, E any](l L) []E {
	res := []E{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

// occurences returns a map counting the amount of each element in the slice
func occurences[L ~[]E, E comparable](l L) map[E]int {
	res := map[E]int{}
	for _, el := range l {
		res[el]++
	}
	return res
}

// intersection performs the intersection operator on two slices
func intersection[L ~[]E, E comparable](l1, l2 L) []E {
	res := []E{}
	for _, el := range l1 {
		if slices.Contains(l2, el) {
			res = append(res, el)
		}
	}
	return res
}

// permutations returns a slice containing all possible permutations of the slice given
func permutations[L ~[]E, E any](l L) [][]E {
	if len(l) <= 1 {
		return [][]E{l}
	}
	res := [][]E{}
	perms := permutations(l[1:])
	for _, perm := range perms {
		for i := 0; i < len(perm)+1; i++ {
			newPerm := append(append([]E{}, perm[:i]...), l[0])
			newPerm = append(newPerm, perm[i:]...)
			res = append(res, newPerm)
		}
	}

	return res
}
