package intervalListIntersections

// 986. Interval List Intersections
// Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
// Return the intersection of these two interval lists.
// (Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b. 
// The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval. 
// For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/interval-list-intersections
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：双指针根据结束值大小更新指针位置
func intervalIntersection(A [][]int, B [][]int) [][]int {
	l, r := 0,0
	ret := make([][]int, 0)
	for l< len(A) && r < len(B) {
		if A[l][1]< B[r][0] {
			l++
			continue
		}
		if B[r][1] < A[l][0] {
			r++
			continue
		}
		ret = append(ret, []int{max(A[l][0], B[r][0]), min(A[l][1], B[r][1])})
		if A[l][1] > B[r][1]{
			r++
		}else {
			l++
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
