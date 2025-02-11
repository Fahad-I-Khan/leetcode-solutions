package main  //  Leetcode 125 - Valid Palindrome

//  s = "A man, a plan, a canal: Panama"
import (
    "fmt"
    "unicode"
)

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
    // Step 1: Filter and normalize the string
    var filtered []rune
    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            filtered = append(filtered, unicode.ToLower(ch))
        }
    }
    
    // Step 2: Check if the filtered string is a palindrome
    left, right := 0, len(filtered)-1
    for left < right {
        if filtered[left] != filtered[right] {
            return false
        }
        left++
        right--
    }
    
    return true
}

func main() {
    // Test case
    s := "A man, a plan, a canal: Panama"
    fmt.Println(isPalindrome(s)) // Output: true
}
