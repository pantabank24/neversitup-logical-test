package project

func Permute(prefix string, str string) []string {
	var result []string
	n := len(str)
	if n == 0 {
		result = append(result, prefix)
		return result
	}

	used := make(map[byte]bool)

	for i := 0; i < n; i++ {
		if used[str[i]] {
			continue
		}

		used[str[i]] = true

		result = append(result, Permute(prefix+string(str[i]), str[:i]+str[i+1:])...)

		used[str[i]] = false
	}

	return result
}
