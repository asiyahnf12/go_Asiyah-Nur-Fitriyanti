package main

import "fmt"

func countOccurrences(arr []string, str string) int {
	count := 0
	for _, val := range arr {
		if val == str {
			count++
		}
	}
	return count
}

func main() {
	// Contoh penggunaan
	arr := []string{"asd", "qwe", "adi", "qwe", "qwe", "asd", "qwe", "asd"}
	str := "asd"

	occurrences := countOccurrences(arr, str)
	fmt.Printf("Jumlah kemunculan %s: %d\n", str, occurrences)
}
