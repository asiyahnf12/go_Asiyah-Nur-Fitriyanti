package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	fmt.Print("Faktor dari ", n, " adalah: ")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
