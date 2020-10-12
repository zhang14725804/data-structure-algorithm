/*
	给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）
	todo：位操作，计算机补码

	三种解法（todo）
*/
func rangeBitwiseAnd(m int, n int) int {
	// 计算机补码（question）
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