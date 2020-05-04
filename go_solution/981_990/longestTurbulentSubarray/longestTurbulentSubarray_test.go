package longestTurbulentSubarray

import "testing"

// 978. Longest Turbulent Subarray
// A subarray A[i], A[i+1], ..., A[j] of A is said to be turbulent if and only if:
// For i <= k < j, A[k] > A[k+1] when k is odd, and A[k] < A[k+1] when k is even;
// OR, for i <= k < j, A[k] > A[k+1] when k is even, and A[k] < A[k+1] when k is odd.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-turbulent-subarray
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxTurbulenceSize(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}
	dp := make([]int, len(A))
	max := 0
	for i :=0; i < len(dp); i++ {
		dp[i] = 1
	}
	prev := 0
	for i:= 1; i< len(A); i++ {
		if prev > 0 && A[i] - A[i-1]<0 {
			dp[i] = dp[i-1] +1
		}else if prev < 0 && A[i] - A[i-1] >0{
			dp[i] = dp[i-1] +1
		}else if A[i] - A[i-1] != 0 {
			dp[i] =2
		}
		prev = A[i] - A[i-1]
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

func TestMaxTurbulenceSize(t *testing.T) {
	input := []int{0,1,1,0,1,0,1,1,0,0}
	t.Log(maxTurbulenceSize(input))
}