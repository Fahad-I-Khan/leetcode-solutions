package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	start := 0
	maxLength := 0

	for end, char := range s {
		if index, exists := charIndexMap[char]; exists && index >= start {
			start = index + 1
		}
		charIndexMap[char] = end
		if currentLength := end - start + 1; currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	fmt.Println("Length of longest substring without repeating characters:", lengthOfLongestSubstring(s))
}