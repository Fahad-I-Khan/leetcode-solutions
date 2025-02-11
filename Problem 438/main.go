package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	pLen := len(p)
	sLen := len(s)
	if pLen > sLen {
		return []int{}
	}

	var result []int
	pCount := make([]int, 26) // Count of characters in p
	sCount := make([]int, 26) // Count of characters in current window of s

	// Count characters in p
	for i := 0; i < pLen; i++ {
		pCount[p[i]-'a']++
	}

	// Count characters in the first window of s
	for i := 0; i < pLen; i++ {
		sCount[s[i]-'a']++
	}

	// Check if the first window is an anagram
	if equal(pCount, sCount) {
		result = append(result, 0)
	}

	// Slide the window over s
	for i := pLen; i < sLen; i++ {
		// Add the new character to the window
		sCount[s[i]-'a']++
		// Remove the character that is sliding out of the window
		sCount[s[i-pLen]-'a']--

		// Check if the current window matches pCount
		if equal(pCount, sCount) {
			result = append(result, i-pLen+1)
		}
	}

	return result
}

// Helper function to compare two frequency arrays
func equal(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	fmt.Println("Anagram starting indices:", result) // Output: [0, 6]
}