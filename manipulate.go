package main

func PermutationsRecursive(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	permutations := []string{}

	for i, j := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := PermutationsRecursive(remaining)
		for _, sp := range subPermutations {
			permutations = append(permutations, string(j)+sp)
		}
	}

	return permutations
}
