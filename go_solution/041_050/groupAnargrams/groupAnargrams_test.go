package groupAnargrams

import (
	"testing"
)

// 049. Group Anagrams
// Given an array of strings, group anagrams together.
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	group := make(map[string][]string, 0)
	for _, s := range strs {
		h := groupHash([]byte(s))
		if _, ok := group[h]; !ok {
			group[h] = []string{s}
		} else {
			group[h] = append(group[h], s)
		}
	}

	rlt := make([][]string, 0)
	for _, v := range group {
		rlt = append(rlt, v)
	}
	return rlt
}

func groupHash(str []byte) string {
	h := []byte("00000000000000000000000000")
	for _, c := range str {
		i := c - 97
		h[i] = h[i] + 1
	}
	return string(h)
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(strs))
}

func TestGroupHash(t *testing.T) {
	t.Log(groupHash([]byte("bob")))
	t.Log(groupHash([]byte("boo")))
}
