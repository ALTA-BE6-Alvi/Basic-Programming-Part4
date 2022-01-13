package main

import "fmt"

func Pow(x, n int) int {
	// your code here
	count := 1
	var hasil int
	hasil = x

	if n > 3 {
		if n % 2 == 0{
			for i := 1; i < n/2; i++ {
				hasil *= x
				count++
			}
			hasil *= hasil
		} else {
			for i := 1; i < n/2; i++ {
				hasil *= x
				count++
			}
			hasil = hasil * x * x * x
		}
	} else {
		for i := 1; i < n; i++ {
			hasil *= x
			count++
		}
	}

	fmt.Print(count)
	fmt.Print("\t")

	return hasil
}

func main() {
	fmt.Println(Pow(2, 3))  // 8
	fmt.Println(Pow(7, 2))  // 49
	fmt.Println(Pow(10, 5)) // 100000
	fmt.Println(Pow(17, 6)) // 24137569
	fmt.Println(Pow(5, 3))  // 125
}
