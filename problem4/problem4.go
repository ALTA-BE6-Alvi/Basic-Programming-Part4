package main

import (
	"fmt"
	"strconv"
)

func MunculSekali(angka string) []int {
	// your code here
	hasil := []int {} 
	duplicate := false

	for i := 0; i < len(angka); i++ {
		for j := 0; j < len(angka); j++ {
			if i == j {
				continue
			} else {
				if angka[i] == angka[j] {
					duplicate = true
				}
			}
		}

		// Check duplicate or not
		if !duplicate {
			number, err := strconv.Atoi(string(angka[i]))

			if err != nil {
				fmt.Println("There is an error")
				fmt.Println(number)
			}

			hasil = append(hasil, number)
		}
		// Reset
		duplicate = false
	}

	return hasil
}

func main() {
	fmt.Println(MunculSekali("1234123"))    // [4]
	fmt.Println(MunculSekali("76523752"))   // [6, 3]
	fmt.Println(MunculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(MunculSekali("1122334455")) // []
	fmt.Println(MunculSekali("0872504"))    // [8 7 2 5 4]
}
