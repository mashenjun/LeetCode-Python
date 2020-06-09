package decodeString

import (
	"strconv"
	"strings"
	"testing"
)

func decodeString(s string) string {
	stack := make([]string, 0)
	idx := 0
	for idx < len(s) {
		if isDigit(s[idx]){
			digits, i := getDigits(s, idx)
			stack = append(stack, digits)
			idx = i
		}else if shouldpush(s[idx]) {
			stack = append(stack, string(s[idx]))
			idx++
		}else {
			tmp := make([]byte, 0)
			idx++
			for stack[len(stack)-1] != "[" {
				tmp = append([]byte(stack[len(stack)-1]),tmp...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			digs, stack := stack[len(stack)-1], stack[:len(stack)-1]
			nums, _ := strconv.Atoi(digs)
			stack = append(stack, strings.Repeat(string(tmp), nums))
		}
	}
	ret := make([]byte, 0)
	for _, s := range stack {
		ret = append(ret, []byte(s)...)
	}
	return string(ret)
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func getDigits(s string, idx int) (string, int) {
	digits := make([]byte, 0)
	for isDigit(s[idx]) {
		digits = append(digits, s[idx])
		idx++
	}
	return string(digits), idx
}

func shouldpush(b byte) bool {
	if (b >= 'a' && b<='z') || (b >= 'A' && b<='Z') || b=='[' {
		return true
	}
	return false
}

func TestDecodeString(t *testing.T) {
	s := "3[a]2[bc]"
	decodeString(s)
}