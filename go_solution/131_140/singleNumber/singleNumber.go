package singleNumber
// 136. Single Number
// 思路： 位操作，通过^操作抵消两个相同的数字
func singleNumber(nums []int) int {
	rlt := 0
	for _, v := range nums {
		rlt ^= v
	}
	return rlt
}
