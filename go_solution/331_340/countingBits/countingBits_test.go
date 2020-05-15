package countingBits

// 338. Counting Bits
// Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/counting-bits
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 动态规划，一维数组，从前往后遍历
// 含义：dp[i] = 数字i二进制形式下1bit的个数
// 递推：dp[i+1] = dp[i] if i+1是偶数
//              = dp[i] + 1 if i+1是奇数
func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i:= 1; i<= num; i++ {
		if i & 1 == 0 {
			dp[i] = dp[i>>1]
		}else {
			dp[i] = dp[i>>1] + 1
		}
	}
	return dp
}