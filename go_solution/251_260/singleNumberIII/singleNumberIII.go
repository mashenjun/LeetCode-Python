package singleNumberIII

// 260. Single Number III
// Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice.
// Find the two elements that appear only once.

func singleNumber(nums []int) []int {
	mask := 0
	for _, v := range nums {
		mask ^= v
	}
	// get the last 1 bit of xor
	mask = mask &(-mask)
	var foo, bar = 0, 0
	for _, v := range nums {
		if v & mask == 0{
			foo ^= v
		}else {
			bar ^= v
		}
	}
	return []int{foo, bar}
}
