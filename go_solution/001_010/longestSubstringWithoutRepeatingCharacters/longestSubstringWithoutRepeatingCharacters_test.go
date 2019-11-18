package longestSubstringWithoutRepeatingCharacters

import "testing"

/*
	3. Longest Substring Without Repeating Characters
	Given a string, find the length of the longest substring without repeating characters.
	IDEA: slicing window
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) <=1 {
		return len(s)
	}
	res, l, store:= 0, 0, map[int32]int{}
	for i, c := range s {
		if _, ok := store[c]; ok {
			if store[c]+1 > l {
				l = store[c]+1
			}
		}
		if i-l+1 > res {
			res = i-l+1
		}
		store[c] = i
	}
	return res
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb") == 3)
	t.Log(lengthOfLongestSubstring("pwwkew") == 3)
	t.Log(lengthOfLongestSubstring("bbbbb") == 1)
	t.Log(lengthOfLongestSubstring("dvdf") == 3)
	t.Log(lengthOfLongestSubstring("tmmzuxt") == 5)
}
