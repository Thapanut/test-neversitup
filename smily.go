package main

func CountSmileyFaces(arr []string) int {
	eyes := []rune{':', ';'}
	nose := []rune{'-', '~'}
	mouth := []rune{')', 'D'}
	count := 0

	for _, face := range arr {
		if len(face) == 2 && contains(eyes, rune(face[0])) && contains(mouth, rune(face[1])) {
			count++
		} else if len(face) == 3 && contains(eyes, rune(face[0])) && contains(nose, rune(face[1])) && contains(mouth, rune(face[2])) {
			count++
		}
	}

	return count
}

func contains(smiles []rune, target rune) bool {
	for _, elem := range smiles {
		if elem == target {
			return true
		}
	}
	return false
}
