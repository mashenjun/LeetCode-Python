package maximumSubarray

import "testing"

// 53. Maximum Subarray

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.

// 思路：动态规划，一维数组
// 含义 dp[i] = 到下标i为止（包括i），元素和最大的连续子串的元素和
// 递推公式 dp[i+1] = max(dp[i-1]+nums[i+1],nums[i+1])
// 注意返回的不是dp的最后一项，而是在计算整个dp数组时出现的最大值
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i< len(nums); i++ {
		if nums[i] > dp[i-1]+ nums[i] {
			dp[i] = nums[i]
		}else {
			dp[i] = dp[i-1]+ nums[i]
		}
		if dp[i] > max {
			max =dp[i]
		}
	}
	return max
}

func Test_MaxSubArray(t *testing.T) {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	t.Log(maxSubArray(input))
}
