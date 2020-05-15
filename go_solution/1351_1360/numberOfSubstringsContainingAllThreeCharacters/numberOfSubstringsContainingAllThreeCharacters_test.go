package numberOfSubstringsContainingAllThreeCharacters

import "testing"

// 1358. Number of Substrings Containing All Three Characters
// Given a string s consisting only of characters a, b and c.
// Return the number of substrings containing at least one occurrence of all these characters a, b and c.
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/number-of-substrings-containing-all-three-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：双指针+滑动窗口；记录[l:r+1]之间abc的出现情况
func numberOfSubstrings(s string) int {
	l, r := 0,0
	ret := 0
	mark := [3]int{}
	for r <= len(s) {
		if mark[0]>0 && mark[1]>0 && mark[2]>0 {
			ret += len(s)-r+1
			mark[s[l]-'a']--
			l++
		}else {
			if r ==len(s) {
				break
			}
			mark[s[r]-'a']++
			r++
		}
	}
	return ret
}

func TestNumberOfSubstring(t *testing.T) {
	t.Log(numberOfSubstrings("abcabc"))
}
