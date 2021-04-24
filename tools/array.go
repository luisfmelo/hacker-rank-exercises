package tools

func CompareUnorderedStringArrays(a, b []string) bool {
	mapElems := map[string]bool{}
	for _, e := range a {
		mapElems[e] = true
	}
	for _, e := range b {
		if _, exists := mapElems[e]; !exists {
			return false
		}
	}

	return true
}
