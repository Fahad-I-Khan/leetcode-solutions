package main

import "fmt"

func characterReplacement(s string, k int) int {
	charCount := make(map[byte]int)
	left := 0
	maxLength := 0
	maxFreq := 0 // This keeps track of the maximum frequency of any character in the current window

	for right := 0; right < len(s); right++ {
		// Add the current character to the count map
		charCount[s[right]]++

		// Update the maximum frequency character in the window
		maxFreq = max(maxFreq, charCount[s[right]])

		// Check if the current window is valid
		windowSize := right - left + 1
		if windowSize-maxFreq > k {
			// If the window is not valid, shrink from the left
			charCount[s[left]]--
			left++
		}

		// Update the maximum length of the valid window
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

// Helper function to find the maximum of two values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "AABABBA"  // This will output 4 
	// s := "AABABBAAABA" // This will output 5.
	k := 1
	result := characterReplacement(s,k)
	fmt.Println("Max length of substring:", result)
}
