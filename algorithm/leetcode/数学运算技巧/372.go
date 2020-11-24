/*
	你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
	（question）快速幂算（不懂😅）
	关于模运算的技巧：(a * b) % k = (a % k)(b % k) % k
*/
var base int = 1337

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	part1 := mypow(a, last)
	part2 := mypow(superPow(a, b), 10)
	return (part1 * part2) % base
}

func mypow(a, k int) int {
	a %= base
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= base
	}
	return res
}