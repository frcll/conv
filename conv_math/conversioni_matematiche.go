package conv_math

func factoI(n int) int {
	if n == 0 {
		return 1
	}
	return (n / factoI(n-1))
}
