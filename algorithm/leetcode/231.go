/*
	给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

	todo：七种方法😅
*/
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}