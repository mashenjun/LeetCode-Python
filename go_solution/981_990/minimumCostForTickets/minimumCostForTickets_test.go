package minimumCostForTickets

import "testing"

// 983. Minimum Cost For Tickets

// 思路：动态规划 一维数组，从后往前遍历
// 含义 dp[i] = 从i天开始，到最后一天的最小花费
// 递推公式： dp[i] = dp[i+1] if i这天不需要旅行
//				   = min(dp[i+1]+costs[0], dp[i+7]+costs[1]+dp[i+30]+costs[2])
func mincostTickets(days []int, costs []int) int {
	if len(days) == 1 {
		return costs[0]
	}
	allDays := make([]bool, days[len(days)-1]+1)
	dp := make([]int, len(allDays)+31) // 关键小技巧, 多申请一个月，方便编程
	for _, v := range days {
		allDays[v] = true
	}

	for i:= len(allDays)-1; i>= 0; i-- {
		if allDays[i] {
			mCost := dp[i+1] + costs[0]
			if i+7 < len(dp) {
				mCost = min(mCost, dp[i+7]+ costs[1])
			}
			if i+30 < len(dp) {
				mCost = min(mCost, dp[i+30] + costs[2])
			}
			dp[i] = mCost
		}else {
			dp[i] = dp[i+1]
		}
	}
	return dp[1]
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func TestMincostTickets(t *testing.T) {
	days := []int{1,2,3,4,6,8,9,10,13,14,16,17,19,21,24,26,27,28,29}
	costs := []int{3, 14, 50}
	t.Log(mincostTickets(days,costs))
}
