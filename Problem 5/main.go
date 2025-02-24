package main

func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    start, end := 0, 0 
    for i := 0; i < len(s); i++ {
        len1, len2 := 0, 0 
        // Only check the even-length center if string length is even
        if len(s)%2 == 0 {
            len1 = expandAroundCenter(s, i, i+1)  // Even-length palindrome
        } else {
            len1 = expandAroundCenter(s, i, i)    // Odd-length palindrome
            len2 = expandAroundCenter(s, i, i+1)  // Even-length palindrome
        }

        // Get the maximum length palindrome found
        maxLen := max(len1, len2)

        // Update start and end if we found a longer palindrome
        if maxLen > end - start {
            start = i - (maxLen - 1)/2
            end = i + maxLen/2
        }
    }
    return s[start:end+1]  // Return the longest palindrome substring
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}