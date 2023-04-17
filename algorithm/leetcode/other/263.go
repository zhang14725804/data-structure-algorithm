var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	// 对n反复除以2,3,5，直到n不包含质因数2,3,5
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}
