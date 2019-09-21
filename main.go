package main

import (
	"fmt"
	"strings"
)

const proverbs = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

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
	lines := strings.Split(proverbs, "\n")

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
