package main

import (
	"fmt"
	"math"
)

func PrimeNumber(number int) bool {
	// your code here
	count := 1

	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number > 2 && number % 2 == 0 {
		return false
	}
	
	var num float64 = float64(number)
    max_div := int(math.Sqrt(num))
	
	i := 3
    for i <= max_div + 1 { 
        if number % i == 0 {
            return false
		}
		i += 2
		count++
	}

	return true
}

func main() {
	fmt.Println(PrimeNumber(1000000007))  // true
	fmt.Println(PrimeNumber(1500450271))  // true
	fmt.Println(PrimeNumber(1000000000))  // false
	fmt.Println(PrimeNumber(10000000019)) // true
	fmt.Println(PrimeNumber(10000000033)) // true
}
