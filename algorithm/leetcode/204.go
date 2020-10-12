/*
	统计所有小于非负整数 n 的质数的数量。
*/

func countPrimes(n int) int {
	count := 0
	// 判断 2 到 sqrt(n) 是否是 n 的因子，如果有一个是，那就表明 n 不是素数。
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 用一个数组表示当前数是否是素数（question）
func countPrimes(n int) int {
	count := 0
	notPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := 2; j*i < n; j++ {
				notPrime[i*j] = true
			}
		}
	}
	return count
}