package main

func climbingStairs(n int) int {
	if n <= 2 {
		return n;
	}
	dp := make([]int n + 1);

	dp[0] = 0 
	dp[1] = 1 
	dp[2] = 2 

	for i = 3; i <= n; n++ {
		dp[i] = dp[i - 1] + dp[1 - 2];
	}
	return dp[n];
}

