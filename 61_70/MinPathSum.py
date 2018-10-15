class Solution:
    def minPathSum(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        n = len(grid)
        m = len(grid[0])
        sums = [[0 for i in range(m)] for j in range(n)]
        for i in reversed(range(m)):
            for j in reversed(range(n)):
                if j == n-1 and i == m-1:
                    sums[j][i] = grid[j][i]
                elif j == n-1:
                    sums[j][i] = grid[j][i] + sums[j][i+1]
                elif i == m - 1:
                    sums[j][i] = grid[j][i] + sums[j + 1][i]
                else:
                    sums[j][i] = min(grid[j][i]+sums[j+1][i], grid[j][i]+sums[j][i+1])

        return sums[0][0]


    def pathSum(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        dp = [0] * len(grid)
        dp[0] = grid[0][0]
        for i in range(1, len(grid)):
            dp[i] = dp[i - 1] + grid[i][0]
        for j in range(1, len(grid[0])):
            for i in range(len(grid)):
                dp[i] = min(dp[i], dp[i - 1]) + grid[i][j] if i > 0 else dp[i] + grid[i][j]
        return dp


grid = [
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
print(Solution().pathSum(grid))