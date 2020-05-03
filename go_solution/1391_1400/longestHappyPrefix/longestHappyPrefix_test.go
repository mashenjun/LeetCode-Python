package longestHappyPrefix

import (
	"fmt"
	"sort"
	"testing"
)

// 1392. Longest Happy Prefix #very important
// A string is called a happy prefix if is a non-empty prefix which is also a suffix (excluding itself).
//Given a string s. Return the longest happy prefix of s .
//Return an empty string if no such prefix exists.

// 该题的解题思路和KMP中构造失效函数（部分匹配表）的思路是一致的
// dp[i] = 截取str[:i];str[:i]所有前缀子串中能和str[:i]的所有后缀子串匹配上的最长的子串的长度
// dp[0] = 0 只有一个字符时，前缀不存在，匹配的前缀子串自然也就不存在
// 递推逻辑:
// 		1. 如果 str[i+1] == str[dp[i]+1] 那么 dp[i+1] = dp[i]+1
//		2. 如果 1 不满足，我们要先找到str[:i]情况下次长的前缀子串和后缀子串匹配上的长度再对最后一位的字符串做判断；而找这个次长的长度有个规律可循：
//				次大匹配必定在最大匹配 str[:dp[i]] 中，所以次大匹配数就是 str[:dp[i]] 的最大匹配数！
// 				https://www.zhihu.com/question/21923021
// 最后结果可以如下表示 str[:dp[len(dp)-1]]
func longestPrefix(s string) string {
	if len(s) <=1 {
		return ""
	}
	dp := make([]int, len(s))
	dp[0] = 0
	for i:= 1; i< len(s); i ++ {
		if s[i] == s[dp[i-1]] {
			dp[i] = dp[i-1]+1
		}else {
			k := dp[i-1]
			for k != 0 && s[k] != s[i] {
				k = dp[k-1]
			}
			if s[i] == s[k] {
				dp[i] = k +1
			}
		}
	}
	fmt.Println(dp)
	return s[:dp[len(dp)-1]]
}

func longestPrefix1(s string) string {
	if len(s) <=1 {
		return ""
	}
	dp := make([]int, len(s))
	dp[0] = 0
	k := 0 // 为什么用共用的k？
	for i:= 1; i< len(s); i ++ {
		for (k != 0 && s[i] != s[k]) {
			k = dp[k-1]
		}
		if s[i] == s[k] {
			k++
		}
		dp[i] = k
	}
	fmt.Println(dp)
	return s[:dp[len(dp)-1]]
}

func Test_LongestPrefix(t *testing.T) {
	input := "acccbaaacccbaac"
	t.Log(longestPrefix(input))
	//t.Log(longestPrefix1(input))
}

func Test_Search(t *testing.T) {
	nums := []int{1,2,3,4,6,7,8,9,10,11,12,13}
	target := 5
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	t.Log(i)
}