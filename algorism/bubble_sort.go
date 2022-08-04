package algorism

func BubbleSortAsc(A []int) []int {
	N := len(A)

	for j := 0; j < N-1; j++ {
		for i := N - 1; i > j; i-- {
			if A[i] < A[i-1] {
				tmp := A[i]
				A[i] = A[i-1]
				A[i-1] = tmp
			}
		}
	}
	return A
}
func BubbleSortDesc(A []int) []int {
	N := len(A)

	for j := 0; j < N-1; j++ {
		for i := N - 1; i > j; i-- {
			if A[i] > A[i-1] {
				tmp := A[i]
				A[i] = A[i-1]
				A[i-1] = tmp
			}
		}
	}
	return A
}
