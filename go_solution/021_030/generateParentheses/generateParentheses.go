package generateParentheses

import "strings"

// 022. Generate Parentheses
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
func generateParenthesis(n int) []string {
	return find(n, n, []string{})
}
// 给定l个左括号，r个右括号，可以找到的组合
func find(l int, r int, str []string) []string {
	if l == 0 && r == 0 {
		return []string{strings.Join(str, "")}
	}
	if l == r {//
		return find(l-1, r, append(str, "("))
	} else if l > 0 { //
		rlt := find(l-1, r, append(str, "("))
		rlt = append(rlt, find(l, r-1, append(str, ")"))...)
		return rlt
	} else {
		return find(l, r-1, append(str, ")"))
	}
}
