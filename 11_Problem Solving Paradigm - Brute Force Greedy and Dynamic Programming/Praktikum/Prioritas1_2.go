package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func main() {
	var numRows int
	fmt.Println("Input numrows : ")
	fmt.Scan(&numRows)
	result := generate(numRows)
	for _, row := range result {
		fmt.Println(row)
	}
}

// 1 -> list 1
// 1 1 -> list 2
// 1 2 1 -> list 3
// 1 3 3 1 -> list 4
