package main

import (
	"fmt"
	"strconv"
)

// My solution
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// Got this from ChatGPT 

// func isPalindrome(n int) bool {
// 	strNum := strconv.Itoa(n)  // Convert int to string
// 	reversed := ""
// 	for i := len(strNum) - 1; i >= 0; i-- {
// 		reversed += string(strNum[i])  // Reverse the string
// 	}
// 	return strNum == reversed  // Check if original string equals reversed string
// }

// This solution is from Leetcode, later see how it is working.

// func isPalindrome(x int) bool {
// 	var n int
// 	j := x
// 	for x > 0 {
// 		n = n*10 + x%10
// 		x /= 10
// 	}
// 	if j == n {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func main()  {
	x := 121
	if isPalindrome(x) {
		fmt.Printf("The int %d is a palindrome\n", x)
	} else {
		fmt.Printf("The int %d is not a palindrome\n", x)
	}
}