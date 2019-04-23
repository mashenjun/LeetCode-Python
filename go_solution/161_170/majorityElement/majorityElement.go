package majorityElement

import "sort"

// 169. Majority Element
// 思路：排序，中间位置的元素一定是众数
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[len(nums)>>1]
}
