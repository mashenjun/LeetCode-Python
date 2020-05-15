package pointOffer

// 找出数组中重复的数字。
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：hashTable查找重复数字 time:O(n) space:O(n)
func findRepeatNumber1(nums []int) int {
	lookup := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := lookup[v]; ok {
			return v
		}
		lookup[v] = struct{}{}
	}
	return -1
}

// 思路：下标定位法 time:O(n) space:O(1)
func findRepeatNumber2(nums []int) int {
	for i, v := range nums {
		if i == v {
			continue
		}
		if i != v {
			if v == nums[v] {
				return v
			}
			nums[i], nums[v] = nums[v], nums[i]
		}
	}
	return -1
}

