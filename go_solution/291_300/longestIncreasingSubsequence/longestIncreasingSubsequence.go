package longestIncreasingSubsequence

/* 300. Longest Increasing Subsequence
Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
*/
// 思路：使用动态规划
// dp[i] 表示nums[:i+1]中上升序列的长度 (包含nums[i]的最长上升序列长度)
// dp[i] = max{dp[j1],...,dp[jn]}+1 where j1,...,jn < i and nums[j]<nums[i]
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	rlt := 1
	// 初始化dp，
	dp := make([]int, len(nums))
	// 一开始，每个元素都能形成一个长度为1的上升序列
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		max := 1
		// 从 dp[0:i-1] 中找到最大值
		for j := 0; j < i; j++ {
			if dp[j] >= max && nums[i] > nums[j] {
				max = dp[j]
				dp[i] = max + 1
			}
		}
		// 得到 dp[i] 再和 rlt 比较，记录最大值
		if dp[i] > rlt {
			rlt = dp[i]
		}
	}
	return rlt
}
