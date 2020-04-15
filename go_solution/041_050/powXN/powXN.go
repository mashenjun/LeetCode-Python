package powXN

// 050. Pow(x, n)
// 思路：牛顿法
func myPow(x float64, n int) float64 {
	// 调整正负号
	if n < 0 {
		x = 1 / x
		n = -n
	}
	rlt := float64(1)
	tmp := x
	for n != 0 {
		if n%2 == 1 {
			rlt *= tmp
		}
		tmp *= tmp
		n /= 2
	}
	return rlt
}
