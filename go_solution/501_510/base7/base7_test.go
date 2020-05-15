package base7

// 504. Base 7
// Given an integer, return its base 7 string representation.

// 思路：取整取余，对取整后的结果继续 取整取余 直达 取整结果为0
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := make([]byte,0)
	if num < 0 {
		sign = append(sign, '-')
		num = - num
	}
	ret := make([]byte, 0)
	for num > 0 {
		m := num / 7
		r := num % 7
		ret = append([]byte{byte(r+'0')}, ret...)
		num  = m
	}
	return string(append(sign, ret...))
}
