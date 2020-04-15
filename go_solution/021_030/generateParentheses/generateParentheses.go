package generateParentheses

import (
	"fmt"
	"strings"
)

// 022. Generate Parentheses
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
func generateParenthesis(n int) []string {
	//return find1(n, n, []string{})
	return find2([]string{}, "", 0,0,n)
}
// 给定l个左括号，r个右括号，可以找到的组合
func find1(l int, r int, str []string) []string {
	fmt.Printf("%v %v %#v\n", l, r, str)
	if l == 0 && r == 0 {
		return []string{strings.Join(str, "")}
	}
	if l == r {
		return find1(l-1, r, append(str, "("))
	} else if l > 0 { //
		rlt := find1(l-1, r, append(str, "("))
		rlt = append(rlt, find1(l, r-1, append(str, ")"))...)
		fmt.Printf("rlt:%#v\n", rlt)
		return rlt
	} else {
		return find1(l, r-1, append(str, ")"))
	}
}

func find2(rlt []string, curr string, open, close, max int) []string{
	// 如果curr使用了所有的括号对，curr就是一种合法的组合
	if len(curr) == max*2 {
		return append(rlt, curr)
	}

	if open< max {
		rlt = find2(rlt, curr+"(", open+1, close, max)
	}
	if close < open {
		rlt = find2(rlt, curr+")", open, close+1, max)
	}
	return rlt
}


