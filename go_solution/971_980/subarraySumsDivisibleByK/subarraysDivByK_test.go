package subarraySumsDivisibleByK

import "testing"

func subarraysDivByK(A []int, K int) int {
	lookup := make(map[int]int, 0)
	ret, sum := 0, 0
	for _, v := range A {
		sum += v
		// sum 可能是负数
		m := (sum % K +K)%K
		ret += lookup[m]
		lookup[m]++
	}
	return ret
}


func TestMod(t *testing.T) {
	t.Log(-10%3)
}
