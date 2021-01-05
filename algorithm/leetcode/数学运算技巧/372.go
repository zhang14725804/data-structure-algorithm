/*
	你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
	问题分解：
	（1）一是如何处理用数组表示的指数
	（2）二是如何得到求模之后的结果
	（3）三是如何高效进行幂运算

	（ question ）
	（1）用递归处理用数组表示的指数
	（2）关于模运算的技巧： (a * b) % k = (a % k)(b % k) % k
	（3） 分奇数偶数两种情况结合递归进行幂运算
*/
var base int = 1337

func superPow(a int, b []int) int {
	// 空数组
	if len(b) == 0 {
		return 1
	}
	// 截取最后一位
	last := b[len(b)-1]
	b = b[:len(b)-1]
	// 用【递归】处理用数组表示的指数
	part1 := mypow(a, last)
	part2 := mypow(superPow(a, b), 10)
	// 每次乘法都要求模
	return (part1 * part2) % base
}

// 计算 a 的 k 次方然后与 base 求模的结果
func mypow1(a, k int) int {
	a %= base // 对因子求模
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= base // 对乘法结果求模
	}
	return res
}

// 高效幂运算
func mypow(a, k int) int {
	// base case
	if k == 0 {
		return 1
	}
	a %= base
	if k%2 == 1 {
		// k是奇数
		return (a * mypow(a, k-1)) % base
	} else {
		// k是偶数
		sub := mypow(a, k/2)
		return (sub * sub) % base
	}
}