package validParentheses

// 020. Valid Parentheses
// 思路：利用栈来保存遍历过的但未成对的括号。
// 在遍历过程，如果访问到的为左括号，直接入栈，如果访问到的是右括号，看是否和栈顶元素匹配。
// 最后判断栈是否为空
func isValid(s string) bool {
	var stack = make([]int, 0)
	for _, c := range s {
		switch c {
		case 40:
			stack = append(stack, 40)
		case 41:
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != 40 {
				return false
			} else { // 括号匹配，出栈
				stack = stack[:len(stack)-1]
			}
		case 91:
			stack = append(stack, 91)
		case 93:
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != 91 {
				return false
			} else { // 括号匹配，出栈
				stack = stack[:len(stack)-1]
			}
		case 123:
			stack = append(stack, 123)
		case 125:
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != 123 {
				return false
			} else { // 括号匹配，出栈
				stack = stack[:len(stack)-1]
			}
		default:
		}
	}
	return len(stack) == 0
}
