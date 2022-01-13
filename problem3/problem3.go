package main

import "fmt"

func JoinArrayRemoveDuplicate(arrayA, arrayB []string) []string {
	// your code here
	answer := arrayA[:]
	duplicate := false

	for i := 0; i < len(arrayB); i++ {
		for j := 0; j < len(arrayA); j++ {
			if arrayB[i] == arrayA[j] {
				duplicate = true
			}
		}
		if !duplicate {
			answer = append(answer, arrayB[i])
		}
		duplicate = false
	}
	
	return answer
}

func main() {
	// Test cases
	fmt.Println(JoinArrayRemoveDuplicate([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}))
	// ["apel", "anggur", "lemon", "leci", "nanas"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}))
	// ["samsung", "apple", "sony", "xiaomi"]

	fmt.Println(JoinArrayRemoveDuplicate([]string{"football", "basketball"}, []string{"basketball", "football"}))
	// [football basketball]
}
