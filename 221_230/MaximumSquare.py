# 221. Maximum Square

class Solution:
    # 动态规划，local为以i,j点为右下角的正方形形成的边长，并维持一个最长值globa
    def maximalSquare(self, matrix):
        """
        :type matrix: List[List[str]]
        :rtype: int
        """
        if not matrix or not matrix[0]:
            return 0

        m = len(matrix)
        n = len(matrix[0])
        local = [[0] * n for i in range(m)]
        globa = 0
        # 设置第一列的值
        for i in range(n):
            if matrix[0][i] == "1":
                local[0][i] = 1
                globa = max(globa, local[0][i])
        # 设置第一行的值
        for i in range(m):
            if matrix[i][0] == "1":
                local[i][0] = 1
                globa = max(globa, local[i][0])

        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j] == "0":
                    local[i][j] = 0
                else:
                    # matrix[i][j]为1时dp的状态转移方程，
                    # 此时要看dp[i-1][j-1], dp[i][j-1]，和dp[i-1][j]这三个位置，我们找其中最小的值，并加上1，就是dp[i][j]的当前值了
                    #  不能有0存在，所以只能取交集，最后再用dp[i][j]的值来更新结果globa的值即可
                    local[i][j] = min(local[i][j-1], local[i-1][j], local[i-1][j-1]) + 1
                globa = max(globa, local[i][j])

        return globa * globa