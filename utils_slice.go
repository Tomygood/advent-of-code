// replace any with desired type
package utils

func contains(l []any, x any) bool {
	for _, el := range l {
		if el == x {
			return true
		}
	}
	return false
}

func noDupes(l []any) []any {
	seen := map[any]bool{}
	res := []any{}
	for _, el := range l {
		if !seen[el] {
			res = append(res, el)
			seen[el] = true
		}
	}

	return res
}

func flatten(lists [][]any) []any {
	res := []any{}
	for _, list := range lists {
		for _, el := range list {
			res = append(res, el)
		}
	}
	return res
}

func union(l1, l2 []any) []any {
	return noDupes(flatten([][]any{l1, l2}))
}

func reverse(l []any) []any {
	res := []any{}
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

func occurences(l []any) map[any]int {
	res := map[any]int{}
	for _, el := range l {
		res[el]++
	}
	return res
}

func intersection(l1, l2 []any) []any {
	res := []any{}
	for _, el := range l1 {
		if contains(l2, el) {
			res = append(res, el)
		}
	}
	return res
}

func Permutations(l []any) [][]any {
	if len(l) <= 1 {
		return [][]any{l}
	}
	res := [][]any{}
	perms := Permutations(l[1:])
	for _, perm := range perms {
		for i := 0; i < len(perm)+1; i++ {
			newPerm := append(append([]any{}, perm[:i]...), l[0])
			newPerm = append(newPerm, perm[i:]...)
			res = append(res, newPerm)
		}
	}

	return res
}
