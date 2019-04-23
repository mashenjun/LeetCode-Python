package slidingWindowMaximum

/* 239. Sliding Window Maximum
Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
*/
//
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	window := make([]int, 0) // store index instead of value
	rlt := make([]int, 0)
	for idx, v := range nums {
		if idx >= k && window[0] <= idx-k {
			window = window[1:] // 窗口向后滑动
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= v {
			window = window[:len(window)-1] // 把窗口内比新进元素小的全部移除
		}
		window = append(window, idx) // 把新的元素append进来
		if idx >= k-1 {
			rlt = append(rlt, nums[window[0]]) // 窗口里的第一个元素永远是窗口里最大的
		}
	}
	return rlt
}

// 方法二：大顶堆
