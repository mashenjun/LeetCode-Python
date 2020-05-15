package binaryNumberWithAlternatingBits

import (
	"testing"
)

// 693. Binary Number with Alternating Bits
// Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/binary-number-with-alternating-bits
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：将n右移一位异或n，检查结果是否全为1。
func hasAlternatingBits(n int) bool {
	m := n >> 1
	xor :=  m ^ n
	return xor&(xor+1) == 0
}

func TestHasAlternatingBits(t *testing.T) {
	t.Log(hasAlternatingBits(5))
}
