package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// your code here
	hasil := []int {}
	i := 0
	element := 0
	count := 0

	for i < len(arr) {
		element = target - arr[i]

		for j := i+1; j < len(arr); j++ {
			if element == arr[j] {
				hasil = append(hasil, i)
				hasil = append(hasil, j)
			}
			count++	
		}
		i++
	}

	return hasil
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
