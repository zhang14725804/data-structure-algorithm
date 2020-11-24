/*
	给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

	三种方法：
		（1）排序，再寻找
		（2）利用等差数列求和公式
		（3）利用异或运算：一个数和它本身做异或运算结果为 0，一个数和 0 做异或运算还是它本身。
*/

// 方法1：利用数学公式
func missingNumber(nums []int) int {
	n := len(nums)
	// 等差数列求和，首项0，公差1，前n项和
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}

// 方法2：异或。两个数组，一个完整的，一个缺失，两个数组异或，结果就是缺失的数字。（相同数字异或等于 0）
func missingNumber(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i] ^ i
	}
	return res
}
