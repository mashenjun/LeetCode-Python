package matrixCellsInDistanceOrder
// 1030. Matrix Cells in Distance Order
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	rlt := [][]int{
		[]int{r0, c0},
	}
	level := 1

	for len(rlt) < R*C {
		deltas := getDelta(level)
		for _, v := range deltas {
			candidates := getCandidates([]int{r0, c0}, v[0], v[1])
			for _, vv := range candidates {
				if vv[0] >= 0 && vv[0] < R && vv[1] >= 0 && vv[1] < C {
					rlt = append(rlt, vv)
				}
			}
		}
		level++
	}
	return rlt
}

func getDelta(level int) [][]int {
	if level == 0 {
		return [][]int{
			[]int{0, 0},
		}
	}
	rlt := make([][]int, 0)
	for i := 0; i <= level; i++ {
		rlt = append(rlt, []int{i, level - i})
	}
	return rlt
}

func getCandidates(o []int, dX, dY int) [][]int {
	if dX == 0 && dY == 0 {
		return [][]int{
			o,
		}
	}
	if dX == 0 {
		return [][]int{
			[]int{o[0], o[1] + dY},
			[]int{o[0], o[1] - dY},
		}
	}
	if dY == 0 {
		return [][]int{
			[]int{o[0] + dX, o[1]},
			[]int{o[0] - dX, o[1]},
		}
	}
	return [][]int{
		[]int{o[0] + dX, o[1] + dY},
		[]int{o[0] + dX, o[1] - dY},
		[]int{o[0] - dX, o[1] + dY},
		[]int{o[0] - dX, o[1] - dY},
	}
}
