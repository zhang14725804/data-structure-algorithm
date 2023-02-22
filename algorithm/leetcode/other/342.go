/*
	给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
	整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x

	你能不使用循环或者递归来完成本题吗？
*/
func isPowerOfFour1(n int) bool {
	if n < 1 {
		return false
	}
	//
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// 位运算(question)
func isPowerOfFour1(n int) bool {
	//先判断是否是 2 的幂:n&(n-1) == 0
	return n > 0 && n&(n-1) == 0 && n&Oxaaaaaaaa == 0
}