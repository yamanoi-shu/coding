package structure

type PriorityQueue struct {
	heap *Heap
}

func NewPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{
		heap: NewHeap(cap),
	}
}

func (q *PriorityQueue) Enqueue(i int) {
	q.heap.Push(i)
}

func (q *PriorityQueue) Dequeue() int {
	return q.heap.Pop()
}
