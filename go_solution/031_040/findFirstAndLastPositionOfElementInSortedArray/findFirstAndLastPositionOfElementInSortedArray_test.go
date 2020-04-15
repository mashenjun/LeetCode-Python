package findFirstAndLastPositionOfElementInSortedArray

import (
	"fmt"
	"testing"
)

// 04. Find First and Last Position of Element in Sorted Array
// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
// Your algorithm's runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1,-1}
	}
	l, r := 0, len(nums)-1
	for l<=r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid-1
		}else if nums[mid]< target {
			l = mid+1
		}else {
			// find the target
			ret := []int{}
			for i:= mid; i>=0; i-- {
				if nums[i] != target{
					ret = append(ret, i+1)
					break
				}
				if i == 0 {
					ret = append(ret, i)
					break
				}
			}
			for i:= mid; i< len(nums); i++ {
				if nums[i] != target{
					ret = append(ret, i-1)
					break
				}
				if i == len(nums)-1 {
					ret = append(ret, i)
					break
				}
			}
			return ret
		}
	}
	return []int{-1,-1}
}

func TestSearchRange(t *testing.T) {
	tt := []struct{
		Nums []int
		Target int
		Expect []int
	}{
		{[]int{1,4}, 4, []int{1,1}},
		{[]int{1}, 1, []int{0,0}},
		{[]int{5,7,7,8,8,10}, 8, []int{3,4}},
		{[]int{5,7,7,8,8,10}, 6, []int{-1,-1}},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			ret := searchRange(tc.Nums, tc.Target)
			if fmt.Sprintf("%v", ret) != fmt.Sprintf("%v", tc.Expect) {
				t.Fatalf("expect:%v, got:%v",tc.Expect, ret)
			}
		})
	}
}