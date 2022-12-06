package conv

import (
	"strconv"
	"strings"
)

func SliceN(n int) []int {
	var sliceN []int
	cont := 0
	for n != 0 {
		sliceN = append(sliceN, n%10)
		n /= 10
		cont++
	}
	return sliceN
}

func SliceNst(n int) []int {
	cifre := strings.Split(strconv.Itoa(int(n)), "")
	s := make([]int, len(cifre))
	for i, r := range cifre {
		s[i], _ = strconv.Atoi(r)
	}
	return s
}
