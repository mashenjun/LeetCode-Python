package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T){
	tt := []struct{
		Val string
		Expect string
	}{
		{"abba", "abba"},
		{"aaaabca", "aaaa"},
		{"ababcdbaa", "aba"},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			ret := longestPalindrome(tc.Val)
			if ret != tc.Expect {
				t.Fatalf("expect:%v, got:%v", tc.Expect, ret)
			}
		})
	}
}
