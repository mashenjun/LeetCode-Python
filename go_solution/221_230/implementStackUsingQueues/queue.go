package implementStackUsingQueues

type Queue struct {
	data []int
}

func InitQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return 0
	}
	rlt := q.data[0]
	q.data = q.data[1:]
	return rlt
}

func (q *Queue) Push(i int) {
	q.data = append(q.data, i)
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Top() int {
	if len(q.data) == 0 {
		return 0
	}
	return q.data[0]
}
