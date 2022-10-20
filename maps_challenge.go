// Maps
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
		Count how many times each word appears in a text. 
		This is a very basic step in many text processing 
		algorithms. To split text to words, use the Fields 
		function from the strings package. Also use the 
		ToLower function from same package to convert all 
		words to lowercase and finally, print the map.
	`
	words := []string(strings.Fields(text))

	wordmap := map[string]int{}
	for _, word := range words {
		punctuation := word[len(word)-1:]
		key := ""
		if punctuation == "." || punctuation == "," {
			key = strings.ToLower(word[:len(word)-2])
		} else {
			key = strings.ToLower(word)
		}
		wordmap[key]++
	}
	fmt.Println(wordmap)
}
