package fourSum

import (
	"testing"
)
import "sort"

/*
	18. 4Sum
 	Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums
	such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
*/

func fourSum(nums []int, target int) [][]int {
	rlt := make([][]int, 0)
	if len(nums)<4 {
		return rlt
	}
	if len(nums) == 4 && nums[0]+nums[1]+nums[2]+nums[3] != target {
		return rlt
	}
	sort.Ints(nums)
	for i:= 0; i<len(nums); i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		for j:= len(nums)-1; j>i; j-- {
			if j< len(nums)-1&& nums[j] == nums[j+1] {
				continue
			}
			lo, hi := i+1, j-1
			for lo < hi {
				if nums[lo] + nums[hi] + nums[i] + nums[j] == target {
					rlt = append(rlt, []int{nums[i], nums[j], nums[lo], nums[hi]})
					lo++
					hi--
					for lo < hi && nums[lo] == nums[lo-1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi+1] {
						hi--
					}
				}else if nums[lo] + nums[hi] + nums[i] + nums[j] < target {
					lo++
				}else {
					hi--
				}
			}
		}
	}
	return rlt
}

func TestFourSum(t *testing.T) {
	t.Log(fourSum([]int{5,5,3,5,1,-5,1,-2}, 4))
}
