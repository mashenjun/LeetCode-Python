package implementQueueUsingStacks

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}
	rlt := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return rlt
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}
