package sorting

import "testing"

func Test_quickSort1(t *testing.T) {
	nums := []int{4,5,3,1,7,1,6,9,}
	quickSort1(nums)
	t.Log(nums)
}

func Test_quickSort2(t *testing.T) {
	nums := []int{4,5,3,1,7,1,6,9,}
	quickSort2(nums)
	t.Log(nums)
}
