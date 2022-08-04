package structure

type Queue struct {
	data []int
	size int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (q *Queue) Enqueue(i int) {
	q.data = append(q.data, i)
	q.size++
}

func (q *Queue) Dequeue() int {
	i := q.data[0]
	q.data = q.data[1:]
	q.size--
	return i
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == cap(q.data)
}

func (q *Queue) Top() int {
	return q.data[0]
}
