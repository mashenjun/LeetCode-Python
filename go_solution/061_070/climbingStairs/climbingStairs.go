package climbingStairs

// 70. Climbing Stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	// init dp
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
