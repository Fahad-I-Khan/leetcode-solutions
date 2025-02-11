// Leetcode 49 - Group Anagrams 

// strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
package main

import (
    "fmt"
    "sort"
)

// Helper function to sort the characters in a string
func sortString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}

// Function to group anagrams
func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)
    
    // Group strings by their sorted version
    for _, str := range strs {
        sortedStr := sortString(str)
        anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
    }
    
    // Convert map values to a slice of slices
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }
    
    return result
}

func main() {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    result := groupAnagrams(strs)
    fmt.Println(result) // Output: [["eat" "tea" "ate"] ["tan" "nat"] ["bat"]]
}
