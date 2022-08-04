package structure

type ArrayList struct {
	data []int
	size int
}

func NewArrayList(cap int) *ArrayList {
	return &ArrayList{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (a *ArrayList) Insert(i int, elem int) {
	a.data = append(a.data, 0)
	for j := i; j < a.size; j++ {
		a.data[j+1] = a.data[j]
	}
	a.data[i] = elem
	a.size++
}

func (a *ArrayList) Erase(i int) {
	a.size--
	for j := i; j < a.size; j++ {
		a.data[j] = a.data[j+1]
	}
	a.data = a.data[:a.size]
}
