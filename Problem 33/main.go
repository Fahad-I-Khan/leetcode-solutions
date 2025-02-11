package main  // Search in Rotated Sorted Array

import "fmt"

func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    
    for left <= right {
        mid := left + (right - left) / 2
        
        // Check if target is at mid
        if nums[mid] == target {
            return mid
        }
        
        // Determine which part is sorted
        if nums[left] <= nums[mid] {
            // Left part is sorted
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right part is sorted
            if target > nums[mid] && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1 // Target not found
}

func main() {
    nums := []int{4, 5, 6, 7, 0, 1, 2}
    target := 0
    result := search(nums, target)
    fmt.Println("Index of target:", result) // Outputs: Index of target: 4
}
