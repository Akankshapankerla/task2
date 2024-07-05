// Create and manipulate maps.
// Example: Write a program to count the frequency of words in a text.

package main

import (
	"fmt"
	"strings"
)

func main4() {
	text := "The natural world is a source of immense beauty and joy, providing us with a healthy environment in which to live."

	wordFreq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,")

		wordFreq[word]++
	}

	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
