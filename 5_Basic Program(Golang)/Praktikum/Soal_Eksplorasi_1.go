package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Meminta pengguna memasukkan kata
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&input)

	// Menghapus karakter-karakter yang tidak perlu dari kata
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))

	// Mengecek apakah kata merupakan palindrome atau bukan
	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrome\n", input)
	} else {
		fmt.Printf("%s bukan palindrome\n", input)
	}
}

// Fungsi untuk memeriksa apakah sebuah string adalah palindrome atau bukan
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
