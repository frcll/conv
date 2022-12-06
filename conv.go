package conv

func sliceN(n int) []int {
	cont, x := 0, n
	for x != 0 {
		x /= 10
		cont++
	}
	sliceN := make([]int, cont)
	for i := 0; i < cont; i++ {
		sliceN = append(sliceN, n%10)
	}
	return sliceN
}
