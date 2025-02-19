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

// My way to solve 

// func isAnagram(s string, t string) bool {
    
//     if len(t) != len(s) {
//         return false
//     }
//     // Frequency map for t
//     tFreq := make(map[byte]int)
//     for i := 0; i < len(t); i++ {
//         tFreq[t[i]]++
//     }

//     right := 0

//     for right < len(s) {
//         currentChar := s[right]
//         if tFreq[currentChar] > 0 {
//             tFreq[currentChar]--
//         } else {
//             return false
//         }
//         right++
        
//     }

//     for _, count := range tFreq {
//         if count > 0 {
//             return false
//         }
//     }
//     return true
// }



// Solution from ChatGPT

// func isAnagram(s string, t string) bool {
//     // If lengths differ, they cannot be anagrams
//     if len(s) != len(t) {
//         return false
//     }

//     // Create a map to count the frequency of characters in string s
//     charCount := make(map[byte]int)

//     // Traverse both strings and update frequencies
//     for i := 0; i < len(s); i++ {
//         charCount[s[i]]++
//         charCount[t[i]]--
//     }

//     // Check if all counts are zero
//     for _, count := range charCount {
//         if count != 0 {
//             return false
//         }
//     }

//     return true
// }

