package main

import "fmt"

func frog(h []int) int {
	dp := make([]int, len(h))
	dp[0] = 0

	for i := 1; i < len(h); i++ {
		cost2 := 0
		// {30, 10, 60, 10, 60, 50}
		cost1 := dp[i-1] + abs(h[i]-h[i-1])
		if i > 1 {
			cost2 = dp[i-2] + abs(h[i]-h[i-2])
		} else {
			cost2 = cost1
		}
		dp[i] = min(cost1, cost2)
	}

	return dp[len(h)-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
