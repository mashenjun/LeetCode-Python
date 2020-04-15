package regularExpressionMatching

import "testing"

// 010. Regular Expression Matching
// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
func isMatch(s string, p string) bool {

	dp := make([][]bool, 0)
	// init the dp[0][:] and fill the dp table
	for i := 0; i < len(s)+1; i++ {
		tmp := make([]bool, len(p)+1)
		dp = append(dp, tmp)
	}
	// dp[0][0] indicates both p and s are empty string
	dp[0][0] = true
	// special case like a*, a*b, a*b*c
	for j:=1;j<len(p)+1;j++ {
		if string(p[j-1]) == "*" {
			dp[0][j] = dp[0][j-2]
		}
	}
	// start deduce
	for i:=1; i< len(s)+1;i++ {
		for j:=1; j< len(p)+1;j++ {
			switch string(p[j-1]) {
			case ".":
				dp[i][j] = dp[i-1][j-1]
			case string(s[i-1]):
				dp[i][j] = dp[i-1][j-1]
			case "*":
				// case: p[j] == "*" with two subcase
				if p[j-2] == s[i-1] || string(p[j-2]) == "." {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}else if p[j-2] != s[i-1]  {
					dp[i][j] = dp[i][j-2]
				}
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
		{"aaa","ab*a*c*a",true},
		{"aa","aaa", false},
		{"ab",".*", true},
		{"aab","c*a*b", true},
		{"aa", "a*", true},
		{"a", ".*", true},
		{"abc", "c*", false},
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
