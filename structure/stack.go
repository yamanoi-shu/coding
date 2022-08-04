package structure

type Stack struct {
	data []int
	size int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
	s.size++
}

func (s *Stack) Pop() int {
	i := s.data[s.size-1]
	s.size--
	s.data = s.data[:s.size-1]
	return i
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == cap(s.data)
}

func (s *Stack) Top() int {
	return s.data[s.size-1]
}
