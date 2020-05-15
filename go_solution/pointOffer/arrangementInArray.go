package pointOffer

// 输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

// 思路：dfs 配合map跳过重复字符串
// todo dfs 非常不熟练

func permutation(s string) []string {
	ret := make([]string, 0)
	b := []byte(s)
	dfs(0, b, &ret)
	return ret
}

func dfs(idx int, b []byte, ret *[]string) {
	if idx == len(b) -1 {
		*ret = append(*ret, string(b))
	}
	lookup := make(map[byte]struct{})
	for i:= idx; i < len(b); i ++ {
		if _, ok := lookup[b[i]]; ok {
			continue
		}
		lookup[b[i]] = struct{}{}
		b[i], b[idx] = b[idx], b[i]
		dfs(idx+1, b, ret)
		b[i], b[idx] = b[idx], b[i]
	}
}
