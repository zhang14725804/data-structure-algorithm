/*
	统计所有【小于】非负整数 n 的质数的数量。
	12.31 面试遇到
*/

// 12.30 我做出来的 超出时间限制
func countPrimes(n int) int {
    if n<2{
        return 0
    }
    var res int
    for i:=2;i<n;i++{
        flag:=0
        for j:=2;j<i;j++{
            if i%j == 0{
				flag++
				break
            }
        }
        if flag==0{
            res++
        }
    }
    return res
}


/*
	方法1：逐个判断优化
	技巧：
		（1）判断是否是质数的时候，只需要判断到Sqrt(n)就好了；
		（2）只要有一个数能被 2～Sqrt(n) 整除就不是质数
*/ 
func countPrimes1(n int) int {
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
	// 开根号
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 方法2：用一个数组表示当前数是否是素数（question）
func countPrimes1(n int) int {
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

// 方法3：在方法3的基础上优化
func countPrimes(n int) int {
	// 初始化数组，初始值true
	notPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		notPrime[i] = true
	}
	// 外循环一次优化，内循环一次优化，妙啊！(question)
	// 由于因子的对称性，其中的 for 循环只需要遍历 [2,sqrt(n)] 就够了。第一处优化
	for i := 2; i*i < n; i++ {
		if notPrime[i] {
			// i 的倍数不可能是素数了。第二处优化
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

