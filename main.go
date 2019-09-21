package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"log"
)

// Return a map where KEYS are all characters used in a sentence and VALUES are counts of the given character
func lettersCounts(sentence string) map[rune]int {
	counts := make(map[rune]int, 0)
	for _, letter := range sentence {
		_, isLetterInMap := counts[letter]
		if !isLetterInMap {
			counts[letter] = strings.Count(sentence, string(letter))
		}
	}
	return counts
}

func main() {
	file, err := os.Open("./proverbs.txt")
    if err != nil {
      log.Fatal(err)
    }
		defer file.Close()
		
	proverbs, err := ioutil.ReadAll(file)
	lines := strings.Split(string(proverbs), "\n")

	for i, proverb := range lines {
		wordCount := len(strings.Fields(proverb))
		// Prints out each proverb on a new line indicating order and word count in each proverb
		// Example: 1. The first proverb has six words. (WC: 6)
		fmt.Printf("%d. %s (WC: %d)\n", i+1, proverb, wordCount)

		for letter, count := range lettersCounts(proverb) {
			// Prints a line below each proverb showing counts of characters used in the given proverb
			fmt.Printf("'%c' = %d, ", letter, count)
		}
		fmt.Printf("\n\n")
	}
}
