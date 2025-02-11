package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
    var result []int
    if len(matrix) == 0 {
        return result
    }
    
    top, bottom := 0, len(matrix) - 1
    left, right := 0, len(matrix[0]) - 1
    
    for top <= bottom && left <= right {
        // Traverse from left to right along the top row
        for col := left; col <= right; col++ {
            result = append(result, matrix[top][col])
        }
        top++
        
        // Traverse from top to bottom along the right column
        for row := top; row <= bottom; row++ {
            result = append(result, matrix[row][right])
        }
        right--
        
        if top <= bottom {
            // Traverse from right to left along the bottom row
            for col := right; col >= left; col-- {
                result = append(result, matrix[bottom][col])
            }
            bottom--
        }
        
        if left <= right {
            // Traverse from bottom to top along the left column
            for row := bottom; row >= top; row-- {
                result = append(result, matrix[row][left])
            }
            left++
        }
    }
    
    return result
}

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    result := spiralOrder(matrix)
    fmt.Println(result) // Output: [1 2 3 6 9 8 7 4 5]
}
