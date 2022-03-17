package unique

func Int(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}
	keys := make(map[int]bool)
	var result []int
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

func String(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	keys := make(map[string]bool)
	var result []string
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
