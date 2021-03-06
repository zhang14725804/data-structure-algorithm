/*
	给定一个整数 n，返回 n! 结果尾数中零的数量。
	问题转化为：n! 最多可以分解出多少个因子 5？

	题型：
	（1）输入一个非负整数 n，请你计算阶乘 n! 的结果末尾有几个 0。
	（2）输入一个非负整数 K，请你计算有多少个 n，满足 n! 的结果末尾恰好有 K 个 0。
*/
func trailingZeroes(n int) int {
	res := 0
	for n > 0 {
		res += n / 5
		n = n / 5 // 取整
	}
	return res
}