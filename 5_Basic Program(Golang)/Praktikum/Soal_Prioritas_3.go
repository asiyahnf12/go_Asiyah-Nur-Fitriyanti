package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	switch {
	case nilai >= 80 && nilai <= 100:
		fmt.Println("Grade: A")
	case nilai >= 65 && nilai <= 79:
		fmt.Println("Grade: B")
	case nilai >= 50 && nilai <= 64:
		fmt.Println("Grade: C")
	case nilai >= 35 && nilai <= 49:
		fmt.Println("Grade: D")
	case nilai >= 0 && nilai <= 34:
		fmt.Println("Grade: E")
	default:
		fmt.Println("Nilai Invalid")
	}
}
