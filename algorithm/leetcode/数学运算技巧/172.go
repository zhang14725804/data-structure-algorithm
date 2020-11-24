/*
	给定一个整数 n，返回 n! 结果尾数中零的数量。
*/
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n = n / 5 // 取整
	}
	return res
}