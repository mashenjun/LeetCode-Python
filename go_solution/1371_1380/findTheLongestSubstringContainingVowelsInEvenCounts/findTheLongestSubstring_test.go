package findTheLongestSubstringContainingVowelsInEvenCounts

import (
	"fmt"
	"testing"
)

// Given the string s, return the size of the longest substring containing each vowel an even number of times.
// That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.


// 思路：首先看到 an even number of times 可以想到利用用位操作中的XOR，一个数字XOR自身 == 0
// 从这点出发，可以用一个5位的bits来表征状态，
// 00000 -> aeiou全部出现过偶数次或者一次都没出现过
// 00001 -> a出现过奇数次，eiou出现过偶数次
// 01001 -> ao出现过奇数次，eiu出现过偶数次
// 10000 -> u出现过奇数次，aeio出现过偶数次
// 11111 -> aeiou全部出现过奇数次
func findTheLongestSubstring(s string) int {
	l, r := 0, 0
	status := 0
	// firstPos用于记录某个状态第一次出现时，下标的位置
	// 也可以用[32]int{}来记录某个状态第一次出现时，下标的位置，但是在处理00000这个状态的时候不如用map来的清晰易懂
	firstPos := make(map[int]int,0)
	ret := 0
	// 00000 这个状态比较特殊，当字符串为空时，也满足这个条件（这个时候没有下标），所以用-1这个值来表式
	firstPos[0] = -1
	for i:= 0; i< len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1<<0
		case 'e':
			status ^= 1<<1
		case 'i':
			status ^= 1<<2
		case 'o':
			status ^= 1<<3
		case 'u':
			status ^= 1<<4
		default:
		}
		if pos, ok := firstPos[status]; !ok {
			firstPos[status] = i
		}else {
			if ret < i-pos {
				ret = i-pos
				l, r = pos, i
			}
		}
	}
	fmt.Printf("l:%v, r:%v\n", l, r)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestFindTheLongestSubstring(t *testing.T) {
	input := "elee"
	t.Log(findTheLongestSubstring(input))
}
