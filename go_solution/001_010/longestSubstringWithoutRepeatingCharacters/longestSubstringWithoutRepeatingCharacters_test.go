package longestSubstringWithoutRepeatingCharacters

import "testing"

/*
	003. Longest Substring Without Repeating Characters
	Given a string, find the length of the longest substring without repeating characters.
	IDEA: slicing window, double pointer
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <=1 {
		return len(s)
	}
	ret, l, store:= 0, 0, map[int32]int{}
	for i, c := range s {
		// 如果检测到有重复字符，更新l到下一个字符的位置
		if _, ok := store[c]; ok {
			if store[c]+1 > l {
				l = store[c]+1
			}
		}
		if i-l+1 > ret {
			ret = i-l+1
		}
		store[c] = i
	}
	return ret
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb") == 3)
	t.Log(lengthOfLongestSubstring("pwwkew") == 3)
	t.Log(lengthOfLongestSubstring("bbbbb") == 1)
	t.Log(lengthOfLongestSubstring("dvdf") == 3)
	t.Log(lengthOfLongestSubstring("tmmzuxt") == 5)
}
