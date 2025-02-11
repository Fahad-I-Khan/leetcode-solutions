package main

// [1,2,3,1]

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0];
	}

	dp := make([]int, len(nums));

	dp[0] = nums[0];
	dp[1] = 0;

	if nums[1] > nums[0] {
		dp[1] = nums[1];
	} else {
		dp[1] = nums[0];
	}

	for i := 2; i < len(nums); i++ {
		if dp[i - 2] + nums[i] > dp[i - 1] {
			dp[i] = dp[i - 2] + nums[i]
		} else {
			dp[i] = dp[i - 1];
		}
	}

	return dp[len(nums) - 1]
}

// Another way to solve 
// This is from ChatGPT

// package main

// func rob(nums []int) int {
//     // Handle edge cases
//     if len(nums) == 0 {
//         return 0
//     }
//     if len(nums) == 1 {
//         return nums[0]
//     }

//     // Initialize variables
//     prev1 := 0 // This will keep track of the maximum amount robbed up to house i-2
//     prev2 := 0 // This will keep track of the maximum amount robbed up to house i-1

//     // Iterate through the houses
//     for _, num := range nums {
//         // Calculate the maximum amount if robbing the current house
//         newPrev := max(prev1 + num, prev2)
//         // Update prev1 and prev2 for the next iteration
//         prev1 = prev2
//         prev2 = newPrev
//     }

//     // The maximum amount that can be robbed is stored in prev2
//     return prev2
// }

// // Helper function to find the maximum of two numbers
// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }
