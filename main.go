package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Permutations:", PermutationsRecursive("a"))
	fmt.Println("Permutations:", PermutationsRecursive("ab"))
	fmt.Println("Permutations:", PermutationsRecursive("abc"))
	fmt.Println("Permutations:", PermutationsRecursive("aabb"))
	fmt.Println("Odd :", FindOddOccurs([]int{7}))                                     //  7
	fmt.Println("Odd :", FindOddOccurs([]int{0}))                                     //  0
	fmt.Println("Odd :", FindOddOccurs([]int{1, 1, 2}))                               //  2
	fmt.Println("Odd :", FindOddOccurs([]int{0, 1, 0, 1, 0}))                         //  0
	fmt.Println("Odd :", FindOddOccurs([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})) //  4
	fmt.Println("Smily", CountSmileyFaces([]string{":)", ";(", ";}", ":-D"}))         // 2
	fmt.Println("Smily", CountSmileyFaces([]string{";D", ":-(", ":-)", ";~)"}))       // 3
	fmt.Println("Smily", CountSmileyFaces([]string{";]", ":[", ";*", ":$", ";-D"}))   // 1
}
