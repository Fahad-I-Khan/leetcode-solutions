package main

func maxSubArray(nums []int) int {
    maxSum := nums[0]  // Initialize maxSum to the first element
    currentSum := nums[0]  // Initialize currentSum to the first element
    
    for i := 1; i < len(nums); i++ {
        // Update currentSum: either extend the current subarray or start a new one
        currentSum = max(nums[i], currentSum + nums[i])
        
        // Update maxSum: keep track of the largest sum
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

// Helper function to find the maximum of two numbers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


