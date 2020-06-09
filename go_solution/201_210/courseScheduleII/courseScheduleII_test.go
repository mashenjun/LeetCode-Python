package courseScheduleII

import "testing"

// 思路：构造入度数组以及邻接表，利用BFS做拓扑排序，
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	inDegs := make([]int, numCourses)

	for _, v := range prerequisites {
		inDegs[v[0]]++
		if edges[v[1]] == nil {
			edges[v[1]] = []int{v[0]}
		}else {
			edges[v[1]] = append( edges[v[1]], v[0])
		}
	}
	queue := make([]int, 0)
	for i, v := range inDegs {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	ret := make([]int, 0)
	var curr int
	for len(queue) != 0 {
		curr, queue = queue[0], queue[1:]
		ret = append(ret, curr)
		for _, v := range edges[curr] {
			inDegs[v]--
			if inDegs[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(ret) != numCourses {
		return []int{}
	}
	return ret
}

func TestFindOrder(t *testing.T) {
	t.Log(findOrder(2, [][]int{{1,0}}))
}
