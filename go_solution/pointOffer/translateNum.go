package pointOffer

import "strconv"

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
// 请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


// 思路：一维动态规划
// 含义：dp[i]=以i为结尾的字符串对应的翻译方法总数
// 递推：dp[i]= dp[i-1] (如果i和i-1无法组成一个数字) or dp[i-1]+dp[i-2] (如果i和i-1可以组成一个数字)
// 返回；dp[-1]
func translateNum(num int) int {
	bs:= strconv.Itoa(num)
	dp := make([]int, len(bs)+1)
	dp[0] = 1
	for i:=1; i < len(dp); i++ {
		dp[i] = dp[i-1]
		if i - 2 >= 0 && bs[i - 2] != '0' && 10 * (bs[i - 2] - '0') + bs[i - 1] - '0' < 26 {
			dp[i] += dp[i - 2]
		}
	}
	return dp[len(bs)]
}