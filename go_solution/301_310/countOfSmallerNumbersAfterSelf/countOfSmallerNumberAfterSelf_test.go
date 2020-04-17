package countOfSmallerNumbersAfterSelf

import "testing"

// 315. Count of Smaller Numbers After Self

type KV struct {
	key int
	val int
}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{0}
	}
	kvs := make([]KV, 0)
	for i, v := range nums {
		kvs = append(kvs, KV{
			key: i,
			val: v,
		})
	}
	ret :=make([]int, len(nums))
	mergeSort(kvs, ret)
	return ret
}

func mergeSort(kvs []KV, counter []int) []KV {
	if len(kvs) <= 1 {
		return kvs
	}
	mid := len(kvs) >> 1
	left := mergeSort(kvs[:mid], counter)
	right := mergeSort(kvs[mid:], counter)
	return merge(left, right, counter)
}

func merge(list1 []KV, list2 []KV, counter []int) []KV {
	ret := make([]KV, 0)
	i, j := 0, 0
	for i < len(list1) || j < len(list2) {
		if j >=len(list2) || (i<len(list1) && list1[i].val <= list2[j].val) {
			counter[list1[i].key] += j
			ret = append(ret, list1[i])
			i++
		} else if i >= len(list1) || (j<len(list2) && list2[j].val < list1[i].val)   {
			ret = append(ret, list2[j])
			j++
		}
	}
	return ret
}

func TestCountSmaller(t *testing.T) {
	t.Log(countSmaller([]int{5,2,6,1}))
}
