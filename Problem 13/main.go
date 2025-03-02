package main

import (
	"fmt"
)

func romanToInt(roman string) int {
	// Map of Roman numerals to their integer values
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize the result variable
	total := 0
	// Iterate through the string
	for i := 0; i < len(roman); i++ {
		// Current Roman numeral character
		current := rune(roman[i])

		// Check if there is a next character and if we need to subtract
		if i+1 < len(roman) {
			next := rune(roman[i+1])
			// If current value is less than the next, subtract the current value
			if romanToIntMap[current] < romanToIntMap[next] {
				total -= romanToIntMap[current]
			} else {
				// Otherwise, add the value
				total += romanToIntMap[current]
			}
		} else {
			// If it's the last character, just add its value
			total += romanToIntMap[current]
		}
	}

	return total
}

func main() {
	// Test the function with Roman numeral string input
	roman := "IX" // Example: IX = 9
	result := romanToInt(roman)
	fmt.Printf("The integer value of Roman numeral %s is %d\n", roman, result)

	roman = "MCMXCIV" // Example: MCMXCIV = 1994
	result = romanToInt(roman)
	fmt.Printf("The integer value of Roman numeral %s is %d\n", roman, result)
}
