package bitwiseANDofNumbersRange

// 201. Bitwise AND of Numbers Range

// Given a range [m, n] where 0 <= m <= n <= 2147483647,
// return the bitwise AND of all numbers in this range, inclusive.

// 思路：find the common prefix of m and n 's binary code.
func rangeBitwiseAnd(m int, n int) int {
	i :=0
	for m!=n {
		m = m >> 1
		n = n >> 1
		i++
	}
	return m<< i
}

