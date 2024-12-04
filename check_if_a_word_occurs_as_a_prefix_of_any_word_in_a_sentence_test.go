package main_test

import "strings"

// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
func isPrefixOfWord(sentence string, searchWord string) int {
	n := 0
	words := strings.Split(sentence, " ")

	for n < len(words) {
		word := words[n]

		if len(word) >= len(searchWord) && word[:len(searchWord)] == searchWord {
			return n + 1
		}

		n++
	}

	return -1
}
