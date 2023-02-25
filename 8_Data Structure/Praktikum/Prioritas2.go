package main

import "fmt"

func cariPasangan(arr []int, target int) []int {
	left, right := 0, len(arr)-1 // inisialisasi pointer left dan right
	for left < right {
		total := arr[left] + arr[right] // hitung jumlah dari dua angka
		if total == target {
			return []int{left, right} // jika jumlah sama dengan target, kembalikan indeks
		} else if total < target {
			left++ // jika total kurang dari target, pindahkan pointer left ke kanan
		} else {
			right-- // jika total lebih dari target, pindahkan pointer right ke kiri
		}
	}
	return nil // jika tidak ditemukan pasangan yang sesuai, kembalikan nilai null
}

func main() {
	fmt.Println(cariPasangan([]int{1, 2, 3, 4, 6}, 6)) //[1, 3]
	fmt.Println(cariPasangan([]int{2, 5, 9, 11}, 11))  //[0, 2]
	fmt.Println(cariPasangan([]int{1, 3, 5, 7}, 12))   //[2, 3]
	fmt.Println(cariPasangan([]int{1, 4, 6, 8}, 10))   //[1, 2]
	fmt.Println(cariPasangan([]int{1, 5, 6, 7}, 6))    //[0, 1]
}
