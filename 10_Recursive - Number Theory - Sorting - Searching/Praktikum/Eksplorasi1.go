package main

import "fmt"

func MaxSequence(arr []int) int {
	// your code here
	max_so_far := 0
	max_ending_here := 0

	for _, num := range arr {
		max_ending_here += num

		if max_ending_here < 0 {
			max_ending_here = 0
		}

		if max_ending_here > max_so_far {
			max_so_far = max_ending_here
		}
	}
	return max_so_far
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
