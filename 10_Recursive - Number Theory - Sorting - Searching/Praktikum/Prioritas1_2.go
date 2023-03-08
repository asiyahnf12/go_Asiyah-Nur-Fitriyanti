package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// kodeku
	freq := make(map[string]int)
	for _, item := range items {
		freq[item]++
	}

	var result []pair
	for item, count := range freq {
		result = append(result, pair{name: item, count: count})
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "js", "js"}))         // golang->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})) // C->1 D->2 B->3 A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))                // football->1 basketball->1 tenis->1
}
