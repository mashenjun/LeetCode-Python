package longestPalindrome

// 005. Longest Palindromic Substring
// idea: can use dp to solve this problem
// dp[i][j] is a boolean value indicates if s[i:j] is a Palindrome string
// dp[i][j] == dp[i-1][j-1] && (s[i]== s[j])
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	maxIdx := [2]int{0, 0} // record i and j for the max result
	// init the dp matrix
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for j := 0; j < len(s); j++ {
		dp[j][j] = true // single character is always a palindromic sub string
		for i := 0; i < j; i++ {
			// todo need explanation 'j-i < 2 || j > 0'
			dp[i][j] = s[j] == s[i] && (j-i < 2 || j > 0 && dp[i+1][j-1])
			if dp[i][j] {
				if j-i > maxIdx[1]-maxIdx[0] {
					maxIdx[0] = i
					maxIdx[1] = j
				}
			}
		}
	}
	return s[maxIdx[0] : maxIdx[1]+1]
}

func longestPalindrome2(s string) string {
	// 初始化dp
	if len(s) == 0 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	l, r := 0, 0
	for j := 0; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if j-i > 1 && dp[i+1][j-1] && s[i] == s[j] { // i与j不邻接
				dp[i][j] = true
			} else if j-i == 1 && s[i] == s[j] { // i与j邻接
				dp[i][j] = true
			}

			if dp[i][j] {
				if r-l < j-i {
					l, r = i, j
				}
			}
		}
	}
	return s[l : r+1]
}
