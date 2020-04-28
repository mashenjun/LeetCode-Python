package uniqueBinarySearchTrees

import "testing"

// 96. Unique Binary Search Trees
// Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
// 使用动态规划, 用一维数组存中间推导结果
// dp[n] = m  含义当有n个数的时候，有m种不同的构造方法
// 初始值: dp[1] = 1; dp[2] = 2
// dp[n+1] = 2*dp[n] + sum(dp[i-1]*dp[n-i]) for i ∈ [2,n-1]
// 返回 dp[len(dp)-1]
func numTrees(n int) int {
	if n <=2 {
		return n
	}
	dp := make([]int,n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i:=3; i<=n; i++ {
		sum := 0
		for left :=0;left < i;left++{
			right := i -1 -left
			if dp[left] == 0 {
				sum += dp[right]
			}else if dp[right] ==0 {
				sum += dp[left]
			}else {
				sum += dp[left]*dp[right]
			}
		}
		dp[i] = sum
	}
	return dp[len(dp)-1]
}

func Test_NumTrees(t *testing.T) {
	input := 4
	t.Log(numTrees(input))

}