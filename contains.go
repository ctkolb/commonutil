package commonutil

func StringMapContains(set map[string]int, item string) bool {
	_, ok := set[item]
	return ok
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
