package algorism

func InsertionSortAsc(A []int) []int {
	N := len(A)
	var j, v int
	for i := 1; i < N; i++ {
		j = i - 1
		v = A[i]
		for j >= 0 && v < A[j] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	return A
}
func InsertionSortDesc(A []int) []int {
	N := len(A)
	var j, v int
	for i := 1; i < N; i++ {
		j = i - 1
		v = A[i]
		for j >= 0 && v > A[j] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	return A
}
