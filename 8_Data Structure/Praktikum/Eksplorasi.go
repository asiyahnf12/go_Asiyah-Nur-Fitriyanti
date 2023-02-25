package main

import (
	"fmt"
	"math"
)

func main() {
	// inisialisasi matriks
	matrix := [3][3]int{
		{1, 2, 3}, // 0
		{4, 5, 6}, // 1
		{9, 8, 9}, // 2
	}

	// hitung jumlah diagonal dari kiri ke kanan
	var diagonal1 int
	for i := 0; i < len(matrix); i++ {
		diagonal1 += matrix[i][i]
	}

	// hitung jumlah diagonal dari kanan ke kiri
	var diagonal2 int
	for i := 0; i < len(matrix); i++ {
		diagonal2 += matrix[i][len(matrix)-1-i]
	}

	// hitung selisih absolut antara jumlah diagonal
	diff := int(math.Abs(float64(diagonal1 - diagonal2)))

	fmt.Println("Matriks:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Printf("Selisih absolut diagonal: %d\n", diff)
}
