package mergeIntervals

// 先根据起始值排序，然后依次计算结束值
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	qSort(intervals, 0, len(intervals)-1)

	ret := [][]int{intervals[0]}
	for i:= 1; i < len(intervals); i++ {
		if ret[len(ret)-1][1] >= intervals[i][0] {
			last := ret[len(ret)-1]
			last[1] = max(last[1], intervals[i][1])
			ret[len(ret)-1] = last
		}else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func qSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(intervals, left, right)
	qSort(intervals, left, pivot-1)
	qSort(intervals, pivot+1, right)
}

func partition(intervals [][]int, left, right int) int {
	pivotVal := intervals[right][0]
	i := left
	for j:= left; j < right; j++ {
		if intervals[j][0] < pivotVal {
			intervals[j], intervals[i] = intervals[i], intervals[j]
			i++
		}
	}
	intervals[i], intervals[right] = intervals[right], intervals[i]
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
