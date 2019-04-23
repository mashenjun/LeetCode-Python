package containsDuplicate

func containsDuplicate(nums []int) bool {
	dict := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := dict[v]; !ok {
			dict[v] = struct{}{}
		}else {
			return true
		}
	}
	return false
}
