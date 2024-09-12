package main

func FindOddOccurs(arr []int) int {

	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}

		if count%2 != 0 {
			return arr[i]
		}
	}
	return -1
}
