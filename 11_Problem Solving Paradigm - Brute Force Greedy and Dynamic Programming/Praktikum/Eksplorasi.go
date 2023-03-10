package main

import "fmt"

func intRoman(num int) string {
	// Array untuk nilai-nilai angka Romawi
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	// Array untuk simbol-simbol angka Romawi
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	// Variabel untuk menyimpan hasil konversi
	roman := ""

	// Looping untuk mengkonversi angka ke angka Romawi
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return roman
}

func main() {
	fmt.Println(intRoman(4))    // Output: IV
	fmt.Println(intRoman(9))    // Output: IX
	fmt.Println(intRoman(23))   // Output: XXIII
	fmt.Println(intRoman(2021)) // Output: MMXXI
	fmt.Println(intRoman(1646)) // Output: MDCXLVI
}
