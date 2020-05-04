package maximumWidthRamp

import "testing"

// 962. Maximum Width Ramp

// Given an array A of integers, a ramp is a tuple (i, j) for which i < j and A[i] <= A[j].  The width of such a ramp is j - i.
// Find the maximum width of a ramp in A.  If one doesn't exist, return 0.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-width-ramp
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：用单调递减栈来解决这一类问题
// 单调栈应用范围
// 求解数组中元素右边第一个比它小的元素的下标，从前往后，构造单调递增栈；
// 求解数组中元素右边第一个比它大的元素的下标，从前往后，构造单调递减栈；
// 求解数组中元素左边第一个比它小的元素的下标，从后往前，构造单调递减栈；
// 求解数组中元素左边第一个比它小的元素的下标，从后往前，构造单调递增栈。

func maxWidthRamp(A []int) int {
	stack := make([]int,0)
	// 构造单调递减栈
	for i:= 0; i < len(A); i++ {
		if len(stack) == 0 || A[stack[len(stack)-1]] > A[i] {
			stack = append(stack, i)
		}
	}

	ret := 0
	j := len(A)-1
	for j > ret {
		for len(stack) != 0 && A[stack[len(stack)-1]]<= A[j] {
			if ret < j - stack[len(stack)-1] {
				ret = j - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
		j--
	}
	return ret
}

func TestMaxWidthRamp(t *testing.T) {
	input := []int{6,0,8,2,1,5}
	t.Log(maxWidthRamp(input))
}
