package maximumProductSubarray

// 152. Maximum Product Subarray
// 思路：使用动态规划，只需要记录到当前访问元素为止，最大的product 和 最小的product(负负得正)
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 考虑负负得正的情况
	max, min := nums[0], nums[0]
	rlt := max
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 { // 正数的情况
			if nums[i]*max > nums[i] {
				max = nums[i] * max
			} else {
				max = nums[i]
			}
			if nums[i]*min < nums[i] {
				min = nums[i] * min
			} else {
				min = nums[i]
			}
		} else { // 负数的情况
			tmp := max
			if nums[i]*min > nums[i] {
				max = nums[i] * min
			} else {
				max = nums[i]
			}
			if nums[i]*tmp < nums[i] {
				min = nums[i] * tmp
			} else {
				min = nums[i]
			}
		}
		if rlt < max {
			rlt = max
		}
	}
	return rlt
}
