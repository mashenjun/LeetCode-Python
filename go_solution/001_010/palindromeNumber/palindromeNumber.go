package palindromeNumber

// 009. Palindrome Number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := make([]int, 0)
	// 将x按十进制拆分放入digits中
	for x > 0 {
		digits = append(digits, x%10)
		x = x / 10
	}
	// 头尾指针夹逼查询
	l, r := 0, len(digits)-1
	for l < r {
		if digits[l] != digits[r] {
			return false
		}
		l++
		r--
	}
	return true
}
