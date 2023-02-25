package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	var counts [10]int
	var unique []int

	// Hitung jumlah kemunculan tiap angka
	for _, digit := range angka {
		counts[digit-'0']++
	}

	// Cari angka-angka yang muncul hanya sekali
	for i, count := range counts {
		if count == 1 {
			unique = append(unique, i)
		}
	}

	return unique
}

func main() {
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("1122334455"))
}
