package contiguousArray

// 525. Contiguous Array

// Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

func findMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	lookup := map[int]int{
		0: -1,
	}

	counter := 0
	max := 0
	for i, v :=range nums {
		if v == 0 {
			counter --
		}else if v == 1 {
			counter++
		}
		if _, ok := lookup[counter]; !ok {
			lookup[counter] = i
		}else {
			if i - lookup[counter] > max {
				max = i - lookup[counter]
			}
		}
	}
	return max
}
