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

// 是否是素数
func isPrime(n int) bool {
	// 开根号
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

// 最高效的方法
func countPrimes(n int) int {
	// 初始化数组
	notPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		notPrime[i] = true
	}
	// 外循环一次优化，内循环一次优化，妙啊！(question)
	// 由于因子的对称性，其中的 for 循环只需要遍历 [2,sqrt(n)] 就够了
	for i := 2; i*i < n; i++ {
		if notPrime[i] {
			// i 的倍数不可能是素数了
			for j := i * i; j < n; j = j + i {
				notPrime[j] = false
			}
		}
	}
	// 统计结果
	count := 0
	for i := 2; i < n; i++ {
		if notPrime[i] {
			count++
		}
	}
	return count
}

