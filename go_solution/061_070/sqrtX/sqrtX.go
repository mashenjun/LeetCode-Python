package sqrtX

// 69. Sqrt(x)
// 思路：二分查找
func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	l, r := 1, x>>1 // sqrt always smaller than x/2
	var rlt int
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
			rlt = mid
		} else {
			r = mid - 1
		}
	}
	return rlt
}
