package jumpGameII

import "testing"

// 045. Jump Game II

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/jump-game-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：从后往前做dp计算
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int,len(nums))
	dp[len(dp)-1] = 0
	for i:= len(dp)-2; i>=0; i-- {
		min := len(nums)+1
		for j:=1; j<= nums[i]; j++ {
			if i+j > len(nums) -1 {
				break
			}
			if dp[i+j] < min {
				min = dp[i+j]
			}
		}
		dp[i] = min +1
	}
	return dp[0]
}

func TestJump(t *testing.T) {
	input := []int{2,3,1,1,4}
	t.Log(jump(input))
}
