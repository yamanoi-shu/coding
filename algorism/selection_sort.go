package algorism

func SelectionSortAsc(A []int) []int {
	N := len(A)
	for i := 0; i < N-1; i++ {
		minI := i
		for j := i; j <= N-1; j++ {
			if A[minI] > A[j] {
				minI = j
			}
		}
		tmp := A[i]
		A[i] = A[minI]
		A[minI] = tmp
	}
	return A
}
func SelectionSortDesc(A []int) []int {
	N := len(A)
	for i := 0; i < N-1; i++ {
		maxI := i
		for j := i; j <= N-1; j++ {
			if A[maxI] < A[j] {
				maxI = j
			}
		}
		tmp := A[i]
		A[i] = A[maxI]
		A[maxI] = tmp
	}
	return A
}
