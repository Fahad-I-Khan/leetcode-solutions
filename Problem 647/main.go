package main

import "fmt"

func countSubstrings(s string) int {
    count := 0

    expandAroundCenter := func(left, right int) {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            count++
            left--
            right++
        }
    }

    for i := 0; i < len(s); i++ {
        expandAroundCenter(i, i)
        expandAroundCenter(i, i+1)
    }
    return count
}

func main()  {
	s := "aaa"
	fmt.Printf("%s string has %d palindrome substring in it.", s, countSubstrings(s))
}