package powXN

// 050. Pow(x, n)
// 思路：牛顿法(快速幂)
func myPow(x float64, n int) float64 {
	// 调整正负号
	if n < 0 {
		n = -n
		x = 1/x
	}
	ret, base := float64(1), x
	for n != 0 {
		if n&1 == 1 {
			ret = base * ret
		}
		base *= base
		n = n >> 1
	}
	return ret
}
