package sorting

// 非递归方式实现
func quickSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	pivot := partition1(nums, left, right)
	partition1(nums, left, pivot-1)
	partition1(nums, pivot+1, right)
}

func partition1(nums []int, left, right int) (pivot int) {
	base := nums[left]
	i, j := left, right
	for i < j {
		// 从右到左，找到第一个比base小的元素
		for i < j && nums[j] > base {
			j--
		}
		// 把找到的元素移动到i的位置
		if i < j {
			nums[i] = nums[j]
			i++
		}
		// 从左到右，找到第一个比base大的元素
		for i < j && nums[i] < base {
			i++
		}
		// 把找到的元素移动到j的位置
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = base
	return i
}

func quickSort2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSortCore(nums, 0, len(nums)-1)
}

func quickSortCore(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition2(nums, left, right)
	quickSortCore(nums , left, pivot-1)
	quickSortCore(nums , pivot+1, right)
}

// 基于'极客时间-数据结构与算法' 递归方式实现
func partition2(nums[]int, left, right int) int {
	i :=left
	pivot := nums[right]
	for j := left; j<= right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
