package main // House Robber 2 

func robLinear(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    
    prev1 := 0 // Maximum amount robbed up to house i-2
    prev2 := 0 // Maximum amount robbed up to house i-1
    
    for _, num := range nums {
        newPrev := max(prev1 + num, prev2)
        prev1 = prev2
        prev2 = newPrev
    }
    
    return prev2
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    // Case 1: Exclude the first house
    case1 := robLinear(nums[1:])
    // Case 2: Exclude the last house
    case2 := robLinear(nums[:len(nums)-1])
    // Return the maximum of both cases
    return max(case1, case2)
}

