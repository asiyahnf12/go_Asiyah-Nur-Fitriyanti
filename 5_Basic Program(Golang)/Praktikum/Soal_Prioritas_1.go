package main

import "fmt"

func main() {
	var sisiatas, sisibawah, tinggi float64
	fmt.Print("Input sisi atas: ")
	fmt.Scanln(&sisiatas)
	fmt.Print("Input sisi bawah: ")
	fmt.Scanln(&sisibawah)
	fmt.Print("Input tinggi: ")
	fmt.Scanln(&tinggi)

	luas := (sisiatas + sisibawah) * tinggi / 2
	fmt.Println("Luas trapesium adalah", luas)
}
