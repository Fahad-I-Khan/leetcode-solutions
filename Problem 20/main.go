package main  // Leetcode 20 - Valid Parentheses

import "fmt"

func isValid(s string) bool {
    // Create a stack to keep track of opening brackets
    stack := []rune{}

    // Map to hold matching pairs
    match := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        switch char {
        case '(', '{', '[':
            // Push opening brackets onto the stack
            stack = append(stack, char)
        case ')', '}', ']':
            // Check if stack is empty or top of stack doesn't match
            if len(stack) == 0 || stack[len(stack)-1] != match[char] {
                return false
            }
            // Pop the stack if it matches
            stack = stack[:len(stack)-1]
        default:
            // Invalid character
            return false
        }
    }

    // If stack is empty, all brackets are matched
    return len(stack) == 0
}

func main() {
    testCases := []string{"](", "()", "({[()]})", "({[)()]}", "({[()]})[]"}
    for _, tc := range testCases {
        fmt.Printf("isValid(%s) = %v\n", tc, isValid(tc))
    }
}


