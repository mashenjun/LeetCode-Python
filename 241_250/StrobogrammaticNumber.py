# 246. Strobogrammatic Number [https://blog.csdn.net/magicbean2/article/details/73997520]

# A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).
# Write a function to determine if a number is strobogrammatic. The number is represented as a string.
# For example, the numbers "69", "88", and "818" are all strobogrammatic.

class Solution:
    def isStrobogrammatic(self, num):
        """
        :type num: str
        :rtype: bool
        """
        dict = {}
        dict['0'] = '0'
        dict['1'] = '1'
        dict['6'] = '9'
        dict['8'] = '8'
        dict['9'] = '6'

        mid = (len(num) + 1) // 2
        for i in range(mid):
            # 假如当前字符不是dict的key，或者后半部分对应的字符不是dict的value，返回false
            if not dict.get(num[i]) or dict[num[i]] != num[len(num)-i-1]:
                return False

        return True


# 247. Strobogrammatic Number II

# A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).
# Find all strobogrammatic numbers that are of length = n.

# For example,
# Given n = 2, return ["11","69","88","96"].

# Hint:
# Try to use recursion and notice that it should recurse with n-2 instead of n-1.

class Solution2:
    # 先列举出外层的字符，即为outer，然后向里面的字符中间添加元素，递归调用函数，用n-2做参数，生成元素，然后运用笛卡儿积
    # 向每个outer的元素中间添加每个lining的元素。
    # 这里有一种情况在于可能存在100001，我们在内圈递归是会出现0000这种，所以在第一次递归最外层没有"00"，后面的递归都加上
    def findStrobogrammatic(self, n):
        """
        :type n: int
        :rtype: list[str]
        """
        if n == 1:
            return ['0', '1', '8']
        if n == 2:
            return ['11', '69', '88', '96']
        res = []
        for elem in self.find(n-2):
            res.append('1' + elem + '1')
            res.append('6' + elem + '9')
            res.append('8' + elem + '8')
            res.append('9' + elem + '6')
        return res
    # find返回头尾都是0的结果
    def find(self, n):
        if n == 1:
            return ['0', '1', '8']
        outer = ['00', '11', '69', '88', '96']
        if n == 2:
            return outer

        res = []
        lining = self.find(n-2)
        for two in outer:
            for elem in lining:
                res.append(two[0] + elem + two[1])

        return res


# 248. Strobogrammatic Number III

# A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).
# Write a function to count the total strobogrammatic numbers that exist in the range of low <= num <= high.

# For example,
# Given low = "50", high = "100", return 3. Because 69, 88, and 96 are three strobogrammatic numbers.

# Note:
# Because the range might be a large number, the low and high numbers are represented as string.

class Solution3:
    def strobogrammaticInRange(self, low, high):
        """
        :type low: str
        :type high: str
        :rtype: int
        """
        m = len(low)
        n = len(high)
        count = 0
        # 计算长度大于m小于n的数的长度，加入到count
        for i in range(m+1, n):
            count += len(self.findStrobogrammatic(i))
        # Python3中的整型是没有限制大小的，可以当作Long类型使用，所以Python3没有Python2的Long类型。
        if m == n:
            for elem in self.findStrobogrammatic(m):
                if int(low) <= int(elem) <= int(high):
                    count += 1
        else:
            for elem in self.findStrobogrammatic(m):
                if int(elem) >= int(low):
                    count += 1
            for elem in self.findStrobogrammatic(n):
                if int(elem) <= int(high):
                    count += 1

        return count


    def findStrobogrammatic(self, n):
        """
        :type n: int
        :rtype: list[str]
        """
        if n == 1:
            return ['0', '1', '8']
        outer = ['11', '69', '88', '96']
        if n == 2:
            return outer

        res = []
        lining = self.findStrobogrammatic(n - 2)
        for two in outer:
            for elem in lining:
                res.append(two[0] + elem + two[1])

        return res

class Solution4:
    # 与第二题解法类似，在递归的过程中与low，high比较，然后更新count的值
    def strobogrammaticInRange(self, low, high):
        """
        :type low: str
        :type high: str
        :rtype: int
        """
        self.res = 0
        self.find(low, high, "")
        self.find(low, high, "0")
        self.find(low, high, "1")
        self.find(low, high, "8")
        return self.res

    def find(self, low, high, curr):
        if len(low) <= len(curr) <= len(high):
            if len(curr) == len(high) and int(curr) > int(high):    # 当前值大于high
                return
            if int(low) <= int(curr) <= int(high) and curr[0] != "0":
                self.res += 1
        if len(high) - len(curr) <= 1:
            return
        self.find(low, high, "0" + curr + "0")
        self.find(low, high, "1" + curr + "1")
        self.find(low, high, "6" + curr + "9")
        self.find(low, high, "8" + curr + "8")
        self.find(low, high, "9" + curr + "6")
