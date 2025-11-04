package main

import (
	"fmt"
	"strings"
)

func word_freq(text string) map[string]int {
	wordFreq := make(map[string]int)

	cleanedText := strings.ToLower(text)

	words := strings.Fields(cleanedText)

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

func main() {
	s := "Bienvenu chez nous"

	Frequencies := word_freq(s)
	fmt.Println("word Frequencies: ")

	for word, count := range Frequencies {
		fmt.Printf(" '%s': %d\n ", word, count)
	}
}
