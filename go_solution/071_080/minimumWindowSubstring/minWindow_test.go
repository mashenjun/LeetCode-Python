package minimumWindowSubstring

import "testing"

func minWindow(s string, t string) string {
	lookup := make(map[byte]int)
	for i:= 0; i<len(t); i++ {
		if _, ok := lookup[t[i]]; !ok {
			lookup[t[i]] = 1
		}else {
			lookup[t[i]]++
		}
	}
	ret := [2]int{-1, len(s)-1}
	// fisrt skip some useless char

	for l, r := 0,0; r<len(s); r++ {
		if _, ok := lookup[s[r]]; ok {
			lookup[s[r]]--
		}
		for valid(lookup) && l<= r {
			if r-l < ret[1]- ret[0] {
				ret[0], ret[1] = l, r
			}
			if _, ok := lookup[s[l]]; ok {
				lookup[s[l]]++
			}
			l++
		}
	}
	if ret[0] == -1 {
		return ""
	}
	return s[ret[0]:ret[1]+1]
}

func valid(lookup map[byte]int) bool{
	for _, v := range lookup {
		if v > 0 {
			return false
		}
	}
	return true
}

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	ss := "ABC"
	t.Log(minWindow(s, ss))
}
