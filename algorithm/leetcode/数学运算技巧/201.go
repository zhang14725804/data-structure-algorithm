/*
	给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）
*/

/*
	方法1：常规操作
*/
func rangeBitwiseAnd(m int, n int) int {
	// 计算机补码 😅😅😅（question）
	if m == INT_MAX {
		return m
	}
	res := m
	for i := m; i <= n; i++ {
		res &= i
		// 计算机补码（question）
		if res == 0 || i == INT_MAX {
			break
		}
	}
	return res
}

/*
	方法1：考虑子问题
	妙啊 😅😅😅
*/
func rangeBitwiseAnd(left int, right int) int {
	// zero 记录我们右移的次数，也就是最低位 0 的个数。
	zeroes := 0
	// m < n，所以至少会有两个数字，所以最低位相与结果一定是 0
	for right > left {
		// 统计最低位 0 的个数
		// 解决了最低位的问题，只需要把 m 和 n 同时右移一位。然后继续按照上边的思路考虑新的最低位的结果即可。
		right >>= 1
		left >>= 1
		zeroes++
	}
	return left << zeroes
}