package pointOffer

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：朴素的遍历，然后考虑下数学性质，缩小下遍历的范围
func findContinuousSequence(target int) [][]int {
	ret := make([][]int, 0)
	for i:= 1; i <= target>>1; i++ {
		sum := 0
		nums := make([]int, 0)
		for j:= i; j<= target; j++ {
			sum += j
			nums = append(nums, j)
			if sum == target {
				ret = append(ret, nums)
				break
			}else if sum > target {
				break
			}
		}
	}
	return ret
}