package structure

type Heap struct {
	data []int
	size int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (h *Heap) Push(elem int) {
	h.data = append(h.data, elem)
	h.size++
	n := h.size - 1
	for n != 0 {
		i := (n - 1) / 2
		if h.data[i] <= elem {
			tmp := h.data[i]
			h.data[i] = h.data[n]
			h.data[n] = tmp
		}
		n = i
	}
}

func (h *Heap) Pop() int {
	out := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.data = h.data[:h.size]
	n := 0
	j := 2*n + 1
	for j < h.size {
		if h.data[j] < h.data[j+1] {
			j++
		}
		if h.data[j] > h.data[n] {
			tmp := h.data[j]
			h.data[j] = h.data[n]
			h.data[n] = tmp
		}
		n = j
		j = 2*n + 1
	}
	return out
}
