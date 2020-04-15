package medianOfTwoSortedArrays

import (
	"testing"
)

// 004. Median of Two Sorted Arrays
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

// idea: 利用二分查找同时修改nums1和nums2上的分隔点，
// 将他们分成LeftPart(nums1)，RightPart(nums1)，LeftPart(nums2)，RightPart(nums2)
// 使其满足:
// 1.#LeftPart(nums1) + #LeftPart(nums2) == #RightPart(nums1) + #RightPart(nums2)
// 2.Max(LeftPart(nums1)) < Min(RightPart(nums2)) && Max(LeftPart(nums1)) < Min(RightPart(nums2))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 在较短的slice上二分查找，可以节省查找次数
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 特殊情况，一个slice为空
	if len(nums1) == 0 {
		if len(nums2)&1 == 1 {
			return float64(nums2[len(nums2)>>1])
		}
		p := len(nums2)>>1
		return float64(nums2[p]+nums2[p-1])/2
	}
	// 二分查找
	// if Max(LeftPart(nums1)) > Min(RightPart(nums2)); then search in left part;
	// if Max(LeftPart(nums2)) > Min(RightPart(nums2)); then search in right part;
	start, end := 0, len(nums1)
	partition1, partition2 := 0, 0
	for start <= end {
		partition1 = (start+end)>>1
		partition2 = ((len(nums1)+len(nums2)+1)>>1) - partition1
		if partition1 >0 && nums1[partition1-1]>nums2[partition2]{
			end = partition1-1
			continue
		}
		if partition1<len(nums1) && nums2[partition2-1]>nums1[partition1] {
			start = partition1+1
			continue
		}
		break
	}
	// find the correct case
	if (len(nums1)+len(nums2))&1 == 1 {
		if partition1 == 0 {
			return float64(nums2[partition2-1])
		}else if partition2 == 0 {
			return float64(nums1[partition1-1])
		}
		return Max(float64(nums1[partition1-1]), float64(nums2[partition2-1]))
	}else {
		l1,l2,r1,r2 := 0.0,0.0,0.0,0.0
		if partition1 == 0 {
			l1, r1 = -10000000000,float64(nums1[partition1])
		} else if partition1 == len(nums1) {
			l1, r1 = float64(nums1[partition1-1]),10000000000
		}else {
			l1, r1 = float64(nums1[partition1-1]),float64(nums1[partition1])
		}
		if partition2 == 0  {
			l2, r2 = -10000000000,float64(nums2[partition2])
		}else if partition2 == len(nums2){
			l2, r2 = float64(nums2[partition2-1]), 10000000000
		}else {
			l2, r2 = float64(nums2[partition2-1]),float64(nums2[partition2])
		}
		return (Max(l1,l2)+Min(r1,r2))/2
	}
}

func Max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Min(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func TestFindMedianSortedArrays(t *testing.T) {
	tt := []struct{
		nums1 []int
		nums2 []int
		expect float64
	}{
		{[]int{1},[]int{2,3,4}, 2.5},
		{[]int{1,2},[]int{3,4}, 2.5},
		{[]int{1,2,3,4}, []int{5,6,7,8},4.5},
		{[]int{5,6,7,8},[]int{1,2,3,4},4.5 },
		{[]int{1,4,5},[]int{2,3,6,7}, 4},
		{[]int{2,3,6,7},[]int{1,4,5}, 4},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			ret := findMedianSortedArrays(tc.nums1, tc.nums2)
			if ret != tc.expect {
				t.Fatalf("expect:%v, got:%v", tc.expect,ret)
			}
		})
	}
}
