package main

func PermutationsRecursive(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	permutations := []string{}
	uniquePerms := make(map[string]bool)

	for i, j := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := PermutationsRecursive(remaining)
		for _, sp := range subPermutations {
			newPerm := string(j) + sp
			if !uniquePerms[newPerm] {
				permutations = append(permutations, newPerm)
				uniquePerms[newPerm] = true
			}
		}
	}

	return permutations
}
