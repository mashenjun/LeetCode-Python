package largestRectangleInHistogram

import "testing"

func largestRectangleArea(heights []int) int {
	ret := 0
	stack := make([]int, 0)
	for i, v := range heights {
		if len(stack) == 0 || v > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for heights[stack[len(stack)-1]] > v {
			top := 0
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			h := heights[top]
			l := 0
			if len(stack) == 0 {
				l = i
			}else {
				l = i - stack[len(stack)-1] - 1
			}
			area := h *l
			if area > ret {
				ret = area
			}
			if len(stack) == 0 {
				break
			}
		}
		stack = append(stack, i)
	}
	for len(stack) != 0 {
		top := 0
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		area, h := 0, heights[top]
		if len(stack) == 0 {
			area = h * len(heights)
		}else {
			area = h * (len(heights)- stack[len(stack)-1]-1)
		}
		if area > ret {
			ret = area
		}
	}
	return ret
}

func Test_LargestRectangleArea(t *testing.T) {
	heights := []int{1,2,2}
	area := largestRectangleArea(heights)
	t.Log(area)
}