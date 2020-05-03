package minimumPathSum

import "testing"

// 064. Minimum Path Sum

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.

// 思路：动态规划，dp是一个二维数组
// 含义：dp[i][j] = 到grid[i][j]这个位置的最小和
// 递推公式：dp[i+1][j+1] = grid[i+1][j+1] + min(dp[i-1][j],dp[i][j-1])
// 注意边界值的情况,即i==0或者j==0的情况
// 结果：dp[len(dp)-1][len(dp[0])-1]
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	ret := minPathSum(grid)
	t.Log(ret)
}
