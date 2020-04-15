package stringToInteger

// 008. String to Integer (atoi)
func myAtoi(str string) int {
	keys := make(map[string]int)
	keys["0"] = 0
	keys["1"] = 1
	keys["2"] = 2
	keys["3"] = 3
	keys["4"] = 4
	keys["5"] = 5
	keys["6"] = 6
	keys["7"] = 7
	keys["8"] = 8
	keys["9"] = 9
	// int32 的最大最小值
	max := (1 << 31) - 1
	min := -1 * (1 << 31)
	rlt := 0
	sign := 1         // 记录正负号
	findSign := false // 记录是否检已经测到符号或数字
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if val, ok := keys[char]; !ok {
			// 如果当前位置的字符不是数字，分别处理4种情况
			if char == " " && findSign == false {
				continue
			} else if char == "-" && findSign == false {
				sign = -1
				findSign = true
			} else if char == "+" && findSign == false {
				sign = 1
				findSign = true
			} else {
				goto END
			}
		} else {
			rlt = 10*rlt + val
			if findSign == false {
				findSign = true
			}
			if sign > 0 && rlt >= max {
				return max
			}
			if sign < 0 && -rlt <= min {
				return min
			}
		}
	}
END:
	return rlt * sign
}
