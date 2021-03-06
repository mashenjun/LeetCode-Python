package topKFrequentElement

/*
347. Top K Frequent Elements
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
*/

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	return topKFrequentWithQS(nums, k)
}

// 思路：使用大顶堆，pop k次，获得的就是前k大的元素
func topKFrequentWithHeap(nums []int, k int) []int {
	fmap := make(map[int]int)
	rlt := make([]int, 0)
	for _, v := range nums {
		if _, ok := fmap[v]; !ok {
			fmap[v] = 0
		}
		fmap[v] = fmap[v] + 1
	}
	h := Heap{
		// capability:k,
		data: make([]Item, 0, k),
	}
	heap.Init(&h)
	for key, value := range fmap {
		item := Item{p: value, v: key}
		heap.Push(&h, item)
	}
	for i := 0; i < k; i++ {
		val := heap.Pop(&h).(Item)
		rlt = append(rlt, val.v)
	}
	return rlt
}

// 思路：使用快排排序
func topKFrequentWithQS(nums []int, k int) []int {
	fmap := make(map[int]int)
	rlt := make([]int, 0)
	for _, v := range nums {
		if _, ok := fmap[v]; !ok {
			fmap[v] = 0
		}
		fmap[v] = fmap[v] + 1
	}
	items := make([]Item, 0)
	for key, value := range fmap {
		item := Item{p: value, v: key}
		items = append(items, item)
	}
	qucikSort(items, 0, len(items)-1)

	for i := 0; i < k; i++ {
		rlt = append(rlt, items[i].v)
	}
	return rlt
}

func qucikSort(items []Item, l int, r int) {
	if l < r {
		pivot := partition(items, l, r)
		qucikSort(items, l, pivot-1)
		qucikSort(items, pivot+1, r)
	}
}

func partition(items []Item, l int, r int) int {
	i, j := l, r
	curr := items[r]
	for i < j {
		for i < j && items[i].p >= curr.p {
			i++
		}
		if i < j {
			items[j] = items[i]
			j--
		}

		for j > i && items[j].p <= curr.p {
			j--
		}
		if j > i {
			items[i] = items[j]
			i++
		}
	}
	items[i] = curr
	return i
}
