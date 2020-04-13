package palindromeNumber

import (
	"strconv"
	"testing"
)


func Test_isPalindrome(t *testing.T) {
	tt := []struct {
		Val int
		Expect bool
	}{
		{101, true},
		{1, true},
		{12, false},
		{22, true},
	}
	for _, tc := range  tt {
		t.Run(strconv.Itoa( tc.Val), func(t *testing.T) {
			if ret := isPalindrome(tc.Val); ret != tc.Expect {
				t.Fatalf("val:%+v, expect:%v, got:%v", tc.Val, tc.Expect, ret)
			}
		})
	}
}
