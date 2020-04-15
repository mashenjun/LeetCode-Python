package wildcardMatching

import "testing"

// 044. Wildcard Matching
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).

func isMatch(s string, p string) bool {
	dp := make([][]bool, 0)
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]bool, len(p)+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for j := 1; j < len(p)+1; j++ {
		if string(p[j-1]) == "*" {
			dp[0][j] = dp[0][j-1]
		}
		//if string(p[j-1]) == "?" || p[j-1] == s[0] {
		//	dp[0][j] = dp[0][j-1]
		//} else if string(p[j-1]) == "*" {
		//	dp[0][j] = dp[0][j-1]
		//}
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}else if string(p[j-1]) == "?" {
				dp[i][j] = dp[i-1][j-1]
			}else if string(p[j-1]) == "*" {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(p)]
}

func TestIsMatch(t *testing.T) {
	tt := []struct {
		Str    string
		Patten string
		Expect bool
	}{
		{"mississippi","m??*ss*?i*pi", false},
		{"aab", "c*a*b", false},
		{"aa", "*", true},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			ret := isMatch(tc.Str, tc.Patten)
			if ret != tc.Expect {
				t.Fatalf("expect:%v, got:%v", tc.Expect, ret)
			}
		})
	}
}
