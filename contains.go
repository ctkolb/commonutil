package commonutil

func StringMapContains(set map[string]int, item string) bool {
	_, ok := set[item]
	return ok
}
