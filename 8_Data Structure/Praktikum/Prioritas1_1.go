package main

import "fmt"

func mergeArray(a []string, b []string) []string {
	// Buat map untuk menyimpan nama yang sudah ada
	seen := make(map[string]bool)

	// Buat slice kosong untuk menampung gabungan array
	result := []string{}

	// Iterasi array pertama dan tambahkan ke slice hasil jika belum ada
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}

	// Iterasi array kedua dan tambahkan ke slice hasil jika belum ada
	for _, val := range b {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Contoh penggunaan
	arr1 := []string{"Alice", "Bob", "Charlie"}
	arr2 := []string{"Bob", "Dave", "Eve"}

	result := mergeArray(arr1, arr2)
	fmt.Println(result) // Output: [Alice Bob Charlie Dave Eve]
}
