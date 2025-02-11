package main  // Find Minimum in Rotated Sorted Array

import "fmt"

func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    
    for left < right {
        mid := left + (right - left) / 2
        
        // Compare mid element with the rightmost element
        if nums[mid] > nums[right] {
            // Minimum is in the right part (excluding mid)
            left = mid + 1
        } else {
            // Minimum is in the left part (including mid)
            right = mid
        }
    }
    
    return nums[left]
}

func main() {
    nums := []int{3, 4, 5, 1, 2}
    result := findMin(nums)
    fmt.Println("Minimum value:", result) // Outputs: Minimum value: 1
}

