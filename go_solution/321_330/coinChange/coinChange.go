package coinChange

/* 322. Coin Change
Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/
// 思路：动态规划，类似求解爬楼梯和斐波那契数列
// dp[i] 表示的是i块钱最少用多少硬币可以凑出来
// dp[j] = min{dp[j-1],dp[j-2],dp[j-5]}
// 最后返回dp[amount-1]
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if len(coins) == 0 {
		return -1
	}
	// init dp
	dp := make([]int, amount)
	for i := 0; i < amount; i++ {
		dp[i] = -1
	}
	// 初始化1的位置
	for _, v := range coins {
		if v <= len(dp) {
			dp[v-1] = 1
		}
	}
	for i := 0; i < amount; i++ {
		min := amount + 1
		for _, v := range coins {
			if i-v >= 0 && dp[i-v] != -1 {
				if dp[i-v] < min {
					min = dp[i-v]
				}
			}
		}
		if min < amount+1 && dp[i] != 1 {
			dp[i] = min + 1
		}
	}
	return dp[amount-1]
}
