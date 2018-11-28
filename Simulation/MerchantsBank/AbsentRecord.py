# 员工考勤记录：
# 给定一个字符串来代表一个员工的考勤纪录，这个纪录仅包含以下两个字符：
# 'A' : Absent，缺勤
# 'P' : Present，到场
# 如果一个员工的考勤纪录中不超过两个'A'(缺勤),那么这个员工会被奖赏。
# 如果你作为一个员工，想在连续N天的考勤周期中获得奖赏，请问有多少种考勤的组合能够满足要求
#
# 输入描述：
# 考勤周期的天数N（正整数）
#
# 输出描述：
# 这N天里能获得奖赏的考勤组合数

# res = C^2_n + C^1_n + C^0_n

while True:
    try:
        n = int(input())
        res = 1 + n + n * (n-1)
        print(res)
    except:
        break