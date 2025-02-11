package main  // Leetcode 242 - Valid Anagram

import (
    "fmt"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    // Create a frequency map for the characters in the first string
    freq := make(map[rune]int)
    
    // Increment counts based on the first string
    for _, char := range s {
        freq[char]++
    }
    
    // Decrement counts based on the second string
    for _, char := range t {
        freq[char]--
        // If a character count goes negative, strings are not anagrams
        if freq[char] < 0 {
            return false
        }
    }
    
    // If all counts are zero, then the strings are anagrams
    return true
}

func main() {
    s := "anagram"
    t := "nagaram"
    fmt.Println(isAnagram(s, t)) // Output: true
}
