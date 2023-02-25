package main

import "fmt"

func power(x int, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		temp := power(x, n/2)
		return temp * temp
	} else {
		temp := power(x, (n-1)/2)
		return temp * temp * x
	}
}

func main() {
	fmt.Println(power(2, 3))  // 8
	fmt.Println(power(5, 3))  // 125
	fmt.Println(power(10, 2)) // 100
	fmt.Println(power(2, 5))  // 32
	fmt.Println(power(7, 3))  // 343
}
